package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/kruftik/nebius-instance-upper/internal/config"
	"github.com/kruftik/nebius-instance-upper/internal/upper"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	parentCtx, parentCancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer parentCancelFn()

	ctx, cancelFn := context.WithCancel(context.Background())

	cfg, err := config.ParseConfig()
	if err != nil {
		fmt.Printf("cannot parse config: %+v", err)
		os.Exit(1)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("cannot init logger: %+v", err)
		os.Exit(1)
	}

	log := logger.Sugar()

	log.Debug("logger initialized")

	log.Info("nebius-instance-upper starting")
	defer log.Debug("nebius-instance-upper completed")

	var creds ycsdk.Credentials

	creds, cfg.FolderID, err = upper.NewCloudConfig(cfg)
	if err != nil {
		log.Errorf("cannot obtain cloud config: %+v", err)
		os.Exit(1)
	}

	log.Infof("folder ID found: %s", cfg.FolderID)

	csdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: creds,
	})
	if err != nil {
		log.Errorf("cannot initialize cloud sdk: %+v", err)
		os.Exit(1)
	}

	log.Info("nebius cloud sdk initialized")

	go func() {
		<-parentCtx.Done()
		cancelFn()
	}()

	upper, err := upper.NewUpperService(
		ctx,
		cfg,
		log,
		csdk.Compute().Instance(),
		csdk.Operation(),
	)

	if err := upper.Run(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Warn(err.Error())
			return
		}

		log.Errorf("cannot run upper: %+v", err)
		os.Exit(1)
	}
}
