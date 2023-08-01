package upper

import (
	"context"
	"fmt"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"strconv"
	"time"
)

const (
	managedInstanceLabel      = "upper-managed-instance"
	managedInstanceGroupLabel = "upper-managed-instance-group"
)

type InstanceSpec struct {
	Instance  *compute.Instance
	StartedAt time.Time
}

func (s *Service) listFolderInstances(ctx context.Context) ([]*compute.Instance, error) {
	var (
		instanceList = make([]*compute.Instance, 0)
		pageToken    string
	)

	for {
		resp, err := s.imgr.List(ctx, &compute.ListInstancesRequest{
			FolderId:  s.cfg.FolderID,
			PageToken: pageToken,
		})
		if err != nil {
			return nil, err
		}

		instanceList = append(instanceList, resp.GetInstances()...)

		pageToken = resp.GetNextPageToken()
		if pageToken == "" {
			break
		}

		s.log.Debug("another instances page is expected")
	}

	return instanceList, nil
}

func (s *Service) filterInstanceListByUpperGroup(instanceList []*compute.Instance, groupName string) []*compute.Instance {
	var (
		instanceGroupList = make([]*compute.Instance, 0, len(instanceList))
	)

	for _, instance := range instanceList {
		v, ok := instance.Labels[managedInstanceLabel]
		if !ok {
			s.log.Debugf("skip instance %s: does not managed by instance-upper", instance.GetId())
			continue
		}

		if bEnabled, err := strconv.ParseBool(v); !bEnabled {
			if err != nil {
				s.log.Error(err, "cannot parse bool value")
				continue
			}

			s.log.Debugf("skip instance %s: disabled", instance.GetId())

			continue
		}

		v = instance.Labels[managedInstanceGroupLabel]
		if v != groupName {
			s.log.Debugf("skip instance %s: group names do not match", instance.GetId())
			continue
		}

		instanceGroupList = append(instanceGroupList, instance)
	}

	return instanceGroupList
}

func (s *Service) getInstanceStartTime(ctx context.Context, instance *compute.Instance) (time.Time, error) {
	var instanceStartTime time.Time

	resp, err := s.imgr.ListOperations(ctx, &compute.ListInstanceOperationsRequest{
		InstanceId: instance.GetId(),
	})
	if err != nil {
		return time.Time{}, fmt.Errorf("cannot request operations list for instance %s: %w", instance.GetId(), err)
	}

	for _, op := range resp.GetOperations() {
		m, err := op.GetMetadata().UnmarshalNew()
		if err != nil {
			return time.Time{}, fmt.Errorf("cannot UnmarshalNew: %w", err)
		}

		switch m.(type) {
		case *compute.StartInstanceMetadata,
			*compute.CreateInstanceMetadata:
		default:
			continue
		}

		if op.CreatedAt.AsTime().After(instanceStartTime) {
			instanceStartTime = op.CreatedAt.AsTime()
		}
	}

	return instanceStartTime, nil
}

func (s *Service) processInstanceList(ctx context.Context, instanceList []*compute.Instance) error {
	groupInstances := s.filterInstanceListByUpperGroup(instanceList, s.cfg.InstanceGroupName)

	var (
		instanceSpecs = make([]InstanceSpec, 0, len(groupInstances))

		bRollingRestartIsNeeded bool

		instanceStartTime time.Time
		err               error
	)

	for _, instance := range groupInstances {
		log := s.log.
			With("instanceID", instance.GetId()).
			With("instanceName", instance.GetName())

		log.Debug("start instance processing")

		if instance.GetStatus() == compute.Instance_STOPPED {
			s.log.Info("instance stopped, starting...")

			if err := s.startInstance(ctx, instance); err != nil {
				log.Warnf("cannot start instance: %+v", err)
			}

			continue
		}

		instanceStartTime, err = s.getInstanceStartTime(ctx, instance)
		if err != nil {
			return fmt.Errorf("cannot determine instance %s (%s) start time: %w", instance.GetId(), instance.GetName(), err)
		}

		log.Debugf("instance started at %s", instanceStartTime)

		if instanceStartTime.Add(s.cfg.RollingRestarts.MinUpDuration).Before(time.Now()) {
			bRollingRestartIsNeeded = true
		}

		instanceSpecs = append(instanceSpecs, InstanceSpec{
			Instance:  instance,
			StartedAt: instanceStartTime,
		})
	}

	if bRollingRestartIsNeeded {
		if err := s.RestartGroup(ctx, instanceSpecs); err != nil {
			return fmt.Errorf("cannot restart nodes group: %w", err)
		}
	}

	return nil
}
