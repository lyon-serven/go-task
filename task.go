package task

import (
	libl "github.com/lyon-serven/go-library/log"
	"github.com/lyon-serven/go-task/scheduler"
)

type Logger = libl.Logger
type Job = scheduler.Job
type Factory = scheduler.Factory
type Registry = scheduler.Registry
type Scheduler = scheduler.Scheduler
type TaskSpec = scheduler.TaskSpec
type BaseJob = scheduler.BaseJob

func NewRegistry() *Registry {
	return scheduler.NewRegistry()
}

func New(logger *Logger, registry *Registry) *Scheduler {
	return scheduler.New(logger, registry)
}

func NewBaseJob(name string, logger *Logger) *BaseJob {
	return scheduler.NewBaseJob(name, logger)
}
