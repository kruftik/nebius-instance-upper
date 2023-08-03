package logging

import (
	"fmt"
	"github.com/kruftik/nebius-instance-upper/internal/config"
	"go.uber.org/zap"
)

func NewLogger(cfg config.AppConfig) (*zap.Logger, error) {
	loggerConfig := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("cannot build logger: %w", err)
	}

	logger.Sugar().Debug("logger initialized")

	return logger, nil
}
