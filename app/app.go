package app

import (
	"strings"

	libl "github.com/lyon-serven/go-library/log"
	task "github.com/lyon-serven/go-task"
	"github.com/lyon-serven/go-task/conf"
)

var (
	Config *conf.Config
	Logger *libl.Logger
)

func Init(configPath string) {
	runtime, err := Load(configPath)
	if err != nil {
		panic("go-task: failed to init app: " + err.Error())
	}
	Config = runtime.Config
	Logger = runtime.Logger
}

func NewJobLogger(name string) (*libl.Logger, error) {
	return NewLogger(&Config.Log, jobLogPath(name))
}

func NewBaseJob(name string) *task.BaseJob {
	return task.NewBaseJob(name, NewJobLogger)
}

func jobLogPath(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ""
	}
	return trimmed + "/"
}
