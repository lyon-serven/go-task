package app

import (
	"strings"

	libl "github.com/lyon-serven/go-library/log"
	task "github.com/lyon-serven/go-task"
	"github.com/lyon-serven/go-task/conf"
	"github.com/lyon-serven/go-task/internal/bootstrap"
)

const defaultConfigPath = bootstrap.DefaultConfigPath

var (
	Config *conf.Config
	Logger *libl.Logger
)

func Init(configPath string) {
	runtime, err := bootstrap.Load(configPath)
	if err != nil {
		panic("go-task: failed to bootstrap runtime: " + err.Error())
	}
	Config = runtime.Config
	Logger = runtime.Logger
}

func NewLogger(name string) (*libl.Logger, error) {
	return bootstrap.NewNamedLogger(&Config.Log, name)
}

func NewBaseJob(name string) *task.BaseJob {
	logger, err := NewLogger(jobLogPath(name))
	if err != nil {
		panic("go-task: failed to init job logger: " + err.Error())
	}
	return task.NewBaseJob(name, logger)
}

func jobLogPath(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ""
	}
	return trimmed + "/"
}
