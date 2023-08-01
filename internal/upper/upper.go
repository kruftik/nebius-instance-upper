package upper

import (
	"context"
	"fmt"
	"github.com/kruftik/nebius-instance-upper/internal/config"
	operation2 "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Upper interface {
}

var (
	_ Upper = (*Service)(nil)
)

type Service struct {
	cfg config.AppConfig
	log *zap.SugaredLogger

	imgr  InstanceManager
	opmgr operation2.OperationServiceClient

	lock bool

	mu sync.Mutex
}

func NewUpperService(_ context.Context, cfg config.AppConfig, log *zap.SugaredLogger, imgr InstanceManager, opmgr operation2.OperationServiceClient) (*Service, error) {
	return &Service{
		cfg: cfg,
		log: log,

		imgr:  imgr,
		opmgr: opmgr,
	}, nil
}

func (s *Service) Process(ctx context.Context) error {
	if s.isRestartProgressing() {
		s.log.Info("skip iteration: rolling restart in progress")
		return nil
	}

	s.log.Info("running process-instances iteration")
	defer s.log.Info("process-instances iteration completed")

	instanceList, err := s.listFolderInstances(ctx)
	if err != nil {
		return fmt.Errorf("cannot load instances: %w", err)
	}

	if err := s.processInstanceList(ctx, instanceList); err != nil {
		return fmt.Errorf("cannot process instances: %w", err)
	}

	return nil
}

func (s *Service) Run(ctx context.Context) error {
	if err := s.Process(ctx); err != nil {
		return err
	}

	ticker := time.NewTicker(s.cfg.InstanceCheckInterval)
	defer ticker.Stop()

	s.log.Info("instance-upper worker started")
	defer s.log.Info("instance-upper worker competed")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := s.Process(ctx); err != nil {
				return err
			}
		}
	}
}
