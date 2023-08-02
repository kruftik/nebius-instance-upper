package upper

import (
	"context"
	"errors"
	"fmt"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"os"
	"os/exec"
	"sort"
	"time"
)

var (
	ErrRollingRestartInProgress = errors.New("rolling-restart in progress")
)

const (
	drainNodeRevertArg = "revert"

	stopNodeTimeout    = 5 * time.Minute
	startNodeTimeout   = 5 * time.Minute
	restartNodeTimeout = 25 * time.Minute

	nodeDrainTimeout       = 5 * time.Minute
	revertNodeDrainTimeout = 5 * time.Second
)

func (s *Service) isRestartProgressing() bool {
	return s.lock
}

func (s *Service) runHookAndWaitForCompletion(ctx context.Context, instance *compute.Instance, command string, args ...string) error {
	cmd := exec.CommandContext(ctx, command, args...)

	home, err := os.UserHomeDir()
	if err != nil {
		s.log.Warnf("cannot detect user home: %+v", err)
		home = "/"
	}

	cmd.Env = append(cmd.Env,
		fmt.Sprintf("HOME=%s", home),
		fmt.Sprintf("NODE_NAME=%s", instance.GetName()),
		fmt.Sprintf("NODE_IP=%s", instance.GetNetworkInterfaces()[0].GetPrimaryV4Address().GetAddress()),
	)

	if kubeConfig := os.Getenv("KUBECONFIG"); kubeConfig != "" {
		s.log.Debugf("inject KUBECONFIG='%s' into script environment", kubeConfig)

		cmd.Env = append(cmd.Env, fmt.Sprintf("KUBECONFIG=%s", kubeConfig))
	} else if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("KUBERNETES_SERVICE_HOST=%s", os.Getenv("KUBERNETES_SERVICE_HOST")),
			fmt.Sprintf("KUBERNETES_SERVICE_PORT=%s", os.Getenv("KUBERNETES_SERVICE_PORT")),
		)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if ctx.Err() != nil {
			return fmt.Errorf("cannot complete %s: %w", command, ctx.Err())
		}

		return fmt.Errorf("cannot complete %s: %w", command, err)
	}

	return nil
}

func (s *Service) restartInstance(ctx context.Context, instance *compute.Instance) error {
	ctx, cancelFn := context.WithTimeout(ctx, restartNodeTimeout)
	defer cancelFn()

	if err := s.stopInstance(ctx, instance); err != nil {
		return fmt.Errorf("cannot stop instance: %w", err)
	}

	if err := s.startInstance(ctx, instance); err != nil {
		return fmt.Errorf("cannot start instance: %w", err)
	}

	return nil
}

func (s *Service) drainNode(ctx context.Context, instance *compute.Instance) error {
	ctx, cancelFn := context.WithTimeout(ctx, nodeDrainTimeout)
	defer cancelFn()

	log := s.log.
		With("instanceID", instance.GetId()).
		With("instanceName", instance.GetName())

	log.Debug("run drain-node script")

	if err := s.runHookAndWaitForCompletion(ctx, instance, s.cfg.DrainNodeScript); err != nil {
		ctx, cancelFn := context.WithTimeout(context.Background(), revertNodeDrainTimeout)
		defer cancelFn()

		if err := s.runHookAndWaitForCompletion(ctx, instance, s.cfg.DrainNodeScript, drainNodeRevertArg); err != nil {
			log.Warnf("cannot revert node drain: %+v", err)
		}

		return fmt.Errorf("drain-node-script failed: %w", err)
	}

	log.Debug("post-start script completed")

	log.Info("instance drained")

	return nil
}

func (s *Service) processInstanceRestart(ctx context.Context, instance *compute.Instance) error {
	log := s.log.
		With("instanceID", instance.GetId()).
		With("instanceName", instance.GetName())

	revertFunc := func(ctx context.Context) {
		log.Info("revert instance drain")

		ctx, cancelFn := context.WithTimeout(ctx, revertNodeDrainTimeout)
		defer cancelFn()

		if err := s.runHookAndWaitForCompletion(ctx, instance, s.cfg.DrainNodeScript, drainNodeRevertArg); err != nil {
			log.Warnf("cannot revert node drain: %+v", err)
		}
	}

	if s.cfg.AlwaysDrain || s.cfg.NodeName == instance.GetName() {
		log.Info("instance drain is required")

		defer revertFunc(context.Background())

		if err := s.drainNode(ctx, instance); err != nil {
			return fmt.Errorf("cannot drain node: %w", err)
		}
	}

	log.Info("restart instance")

	if err := s.restartInstance(ctx, instance); err != nil {
		return fmt.Errorf("cannot restart instance: %w", err)
	}

	log.Info("instance restarted")

	return nil
}

func (s *Service) RestartGroup(ctx context.Context, instanceSpecsList []InstanceSpec) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.lock {
		return ErrRollingRestartInProgress
	}

	s.lock = true
	defer func() {
		s.lock = false
	}()

	sort.SliceStable(instanceSpecsList, func(i, j int) bool {
		return !instanceSpecsList[i].StartedAt.Before(instanceSpecsList[j].StartedAt)
	})

	for _, spec := range instanceSpecsList {
		if spec.StartedAt.Add(s.cfg.RollingRestarts.MinUpDuration).After(time.Now()) {
			s.log.Debug("instance %s (%s) does not need to be restarted", spec.Instance.GetId(), spec.Instance.GetName())
			continue
		}

		if err := s.processInstanceRestart(ctx, spec.Instance); err != nil {
			return fmt.Errorf("cannot process instance %s (%s): %w", spec.Instance.GetId(), spec.Instance.GetName(), err)
		}
	}

	return nil
}
