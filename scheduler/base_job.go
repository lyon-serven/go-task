package scheduler

import (
	"fmt"
	"sync"
	"time"

	libl "github.com/lyon-serven/go-library/log"
	"go.uber.org/zap"
)

type BaseJob struct {
	name   string
	Args   map[string]any
	Logger *libl.Logger
	mu     sync.Mutex
}

func NewBaseJob(name string, logger *libl.Logger) *BaseJob {
	if logger == nil {
		panic("scheduler: logger is required")
	}

	return &BaseJob{
		name:   name,
		Logger: logger,
		Args:   make(map[string]any),
	}
}

func (b *BaseJob) Name() string { return b.name }

func (b *BaseJob) SetArgs(args map[string]any) { b.Args = args }

func (b *BaseJob) Run(job Job) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	defer func() {
		if r := recover(); r != nil {
			b.Logger.Error("job panic", zap.String("job", b.name), zap.Any("panic", r))
		}
	}()

	start := time.Now()
	b.Logger.Info("job started", zap.String("job", b.name))

	err := job.Exec()
	elapsed := time.Since(start)

	if err != nil {
		b.Logger.Error("job failed",
			zap.String("job", b.name),
			zap.Error(err),
			zap.Float64("elapsed_s", elapsed.Seconds()),
		)
	} else {
		b.Logger.Info("job succeeded",
			zap.String("job", b.name),
			zap.Float64("elapsed_s", elapsed.Seconds()),
		)
	}
	return err
}

func (b *BaseJob) Exec() error {
	b.Logger.Info(fmt.Sprintf("job %s: no-op exec at %s", b.name, time.Now().Format(time.DateTime)))
	return nil
}
