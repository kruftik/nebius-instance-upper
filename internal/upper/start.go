package upper

import (
	"context"
	"fmt"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-sdk/operation"
)

func (s *Service) startInstance(ctx context.Context, instance *compute.Instance) error {
	log := s.log.
		With("instanceID", instance.GetId()).
		With("instanceName", instance.GetName())

	startCtx, cancelFn := context.WithTimeout(ctx, startNodeTimeout)
	defer cancelFn()

	opProto, err := s.imgr.Start(startCtx, &compute.StartInstanceRequest{
		InstanceId: instance.GetId(),
	})
	if err != nil {
		return fmt.Errorf("cannot send start-instance-request: %w", err)
	}

	log.Debugf("start request sent, waiting")

	op := operation.New(s.opmgr, opProto)

	if err := op.Wait(startCtx); err != nil {
		return fmt.Errorf("start operation failed: %w", err)
	}

	log.Debug("run post-start script")

	postStartCtx, cancelFn := context.WithTimeout(ctx, startNodeTimeout)
	defer cancelFn()

	if err := s.runHookAndWaitForCompletion(postStartCtx, instance, s.cfg.PostStartScript); err != nil {
		return fmt.Errorf("post-start-script failed: %w", err)
	}

	log.Debug("post-start script completed")

	log.Infof("start completed")

	return nil
}
