package scheduler

import (
	"fmt"
	"sync"
	"time"

	libl "github.com/lyon-serven/go-library/log"
	"go.uber.org/zap"
)

type LoggerFactory func(name string) (*libl.Logger, error)

type BaseJob struct {
	name          string
	Args          map[string]any
	Logger        *libl.Logger
	loggerFactory LoggerFactory
	mu            sync.Mutex
	loggerMu      sync.Mutex
}

func NewBaseJob(name string, factory LoggerFactory) *BaseJob {
	if factory == nil {
		panic("scheduler: logger factory is required")
	}

	return &BaseJob{
		name:          name,
		loggerFactory: factory,
		Args:          make(map[string]any),
	}
}

func (b *BaseJob) Name() string { return b.name }

func (b *BaseJob) Setup() {
	if b.Logger != nil {
		return
	}

	b.loggerMu.Lock()
	defer b.loggerMu.Unlock()

	if b.Logger != nil {
		return
	}
	if b.loggerFactory == nil {
		panic("scheduler: logger is required")
	}

	logger, err := b.loggerFactory(b.name)
	if err != nil {
		panic("scheduler: failed to init job logger: " + err.Error())
	}
	if logger == nil {
		panic("scheduler: logger factory returned nil")
	}
	b.Logger = logger
}

func (b *BaseJob) SetArgs(args map[string]any) { b.Args = args }

func (b *BaseJob) Run(job Job) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Setup()

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
	b.Setup()
	b.Logger.Info(fmt.Sprintf("job %s: no-op exec at %s", b.name, time.Now().Format(time.DateTime)))
	return nil
}
