package config

import (
	"fmt"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

type RollingRestartsConfig struct {
	MinUpDuration time.Duration
}

type AppConfig struct {
	IAMServiceAccountKeyFile string `env:"SERVICE_ACCOUNT_KEY_FILE" required:"true"`
	IAMServiceAccountKey     []byte

	FolderID string `long:"folder-id" env:"FOLDER_ID"`

	InstanceGroupName     string        `long:"instance-group-name" env:"INSTANCE_GROUP_NAME" default:""`
	InstanceCheckInterval time.Duration `long:"instance-check-interval" env:"INSTANCE_CHECK_INTERVAL" default:"15s"`

	PreStopScript   string `long:"pre-stop-script" env:"PRE_STOP_SCRIPT"`
	PostStartScript string `long:"post-start-script" env:"POST_START_SCRIPT"`
	DrainNodeScript string `long:"drain-node-script" env:"DRAIN_NODE_SCRIPT"`

	AlwaysDrain bool `long:"always-drain-nodes" env:"NODES_ALWAYS_DRAIN"`

	LogLevel int `short:"v" long:"v" env:"LOG_LEVEL"`

	NodeName string `long:"node-name" env:"NODE_NAME"`

	RollingRestarts RollingRestartsConfig
}

const (
	envServiceAccountJSON = "SERVICE_ACCOUNT_JSON"
)

func ParseConfig() (AppConfig, error) {
	var cfg AppConfig

	_, err := flags.Parse(&cfg)
	if err != nil {
		return AppConfig{}, fmt.Errorf("cannot parse config: %w", err)
	}

	saJSONString := os.Getenv(envServiceAccountJSON)
	if saJSONString == "" {
		cfg.IAMServiceAccountKey, err = os.ReadFile(cfg.IAMServiceAccountKeyFile)
		if err != nil {
			return AppConfig{}, fmt.Errorf("invalid service-account-key-file config")
		}
	} else {
		cfg.IAMServiceAccountKey = []byte(saJSONString)
	}

	cfg.RollingRestarts = RollingRestartsConfig{
		MinUpDuration: 12 * time.Hour,
	}

	if _, err = os.Stat(cfg.PreStopScript); err != nil {
		return AppConfig{}, fmt.Errorf("invalid pre-stop-script: %w", err)
	}
	if _, err = os.Stat(cfg.PostStartScript); err != nil {
		return AppConfig{}, fmt.Errorf("invalid post-start-script: %w", err)
	}
	if _, err = os.Stat(cfg.DrainNodeScript); err != nil {
		return AppConfig{}, fmt.Errorf("invalid drain-node-script: %w", err)
	}

	if cfg.NodeName == "" {
		if cfg.NodeName, err = os.Hostname(); err != nil {
			return AppConfig{}, fmt.Errorf("cannot get hostname: %w", err)
		}
	}

	return cfg, nil
}
