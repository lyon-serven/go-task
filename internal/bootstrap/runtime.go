package bootstrap

import (
	"time"

	libc "github.com/lyon-serven/go-library/config"
	libl "github.com/lyon-serven/go-library/log"
	"github.com/lyon-serven/go-task/conf"
)

const DefaultConfigPath = "configs/config.yaml"

type Runtime struct {
	Config *conf.Config
	Logger *libl.Logger
}

func Load(configPath string) (*Runtime, error) {
	if configPath == "" {
		configPath = DefaultConfigPath
	}

	cfg := &conf.Config{}
	if err := libc.LoadYAMLConfig(configPath, cfg); err != nil {
		return nil, err
	}

	logger, err := NewLogger(&cfg.Log, "")
	if err != nil {
		return nil, err
	}

	logger.Info("go-task initialized")
	return &Runtime{Config: cfg, Logger: logger}, nil
}

func NewNamedLogger(cfg *conf.Log, name string) (*libl.Logger, error) {
	return NewLogger(cfg, name)
}

func NewLogger(cfg *conf.Log, name string) (*libl.Logger, error) {
	if cfg == nil {
		cfg = &conf.Log{}
	}

	rotationTime := 24 * time.Hour
	if cfg.RotationTime != "" {
		if d, err := time.ParseDuration(cfg.RotationTime); err == nil && d > 0 {
			rotationTime = d
		}
	}

	fileAge := 4320 * time.Hour
	if cfg.FileAge != "" {
		if d, err := time.ParseDuration(cfg.FileAge); err == nil && d > 0 {
			fileAge = d
		}
	}

	fileName := cfg.FileName
	if fileName == "" {
		fileName = "task.log"
	}

	filePattern := cfg.FilePattern
	if filePattern == "" {
		filePattern = "/%Y_%m/%d/task_%H.log"
	}

	environment := cfg.Environment
	if environment == "" {
		environment = "production"
	}

	systemName := cfg.SystemName
	if systemName == "" {
		systemName = "go-task"
	}

	logToFile := cfg.FilePath != ""
	lcf := &libl.LogConfig{
		LogToStdout:  !logToFile,
		LogToFile:    logToFile,
		Level:        cfg.Level,
		Path:         cfg.FilePath + name,
		FileName:     fileName,
		FilePattern:  filePattern,
		RotationTime: rotationTime,
		FileAge:      fileAge,
		Environment:  environment,
		SystemName:   systemName,
	}
	return libl.NewLogger(lcf)
}
