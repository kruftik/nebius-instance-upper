package upper

import (
	"context"
	"fmt"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-sdk/operation"
)

func (s *Service) stopInstance(ctx context.Context, instance *compute.Instance) error {
	log := s.log.
		With("instanceID", instance.GetId()).
		With("instanceName", instance.GetName())

	preStopCtx, cancelFn := context.WithTimeout(ctx, stopNodeTimeout)
	defer cancelFn()

	log.Debug("run pre-stop script")
	if err := s.runHookAndWaitForCompletion(preStopCtx, instance, s.cfg.PreStopScript); err != nil {
		return fmt.Errorf("pre-stop-script failed: %w", err)
	}
	log.Debug("pre-stop script completed")

	stopCtx, cancelFn := context.WithTimeout(ctx, stopNodeTimeout)
	defer cancelFn()

	opProto, err := s.imgr.Stop(stopCtx, &compute.StopInstanceRequest{
		InstanceId: instance.GetId(),
	})
	if err != nil {
		return fmt.Errorf("cannot restart instance %s: %w", instance.GetId(), err)
	}

	log.Debugf("stop request sent, waiting")

	op := operation.New(s.opmgr, opProto)

	if err := op.Wait(stopCtx); err != nil {
		return fmt.Errorf("stop operation failed: %w", err)
	}

	log.Debugf("shutdown completed")

	return nil
}
