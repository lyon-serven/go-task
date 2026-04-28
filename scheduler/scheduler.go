package scheduler

import (
	"errors"
	"time"

	libl "github.com/lyon-serven/go-library/log"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
)

type scheduledJob struct {
	spec TaskSpec
	job  Job
}

type Scheduler struct {
	inner    *gocron.Scheduler
	logger   *libl.Logger
	registry *Registry
	jobs     []scheduledJob
}

func New(logger *libl.Logger, registry *Registry) *Scheduler {
	if logger == nil {
		panic("scheduler: logger is required")
	}
	if registry == nil {
		registry = NewRegistry()
	}
	return &Scheduler{
		inner:    gocron.NewScheduler(time.Local),
		logger:   logger,
		registry: registry,
	}
}

func (s *Scheduler) Load(specs []TaskSpec) error {
	jobs := make([]scheduledJob, 0, len(specs))
	for _, spec := range specs {
		if spec.Job == "" {
			return errors.New("scheduler: task job is required")
		}
		if spec.Cron == "" {
			return errors.New("scheduler: task cron is required")
		}
		if spec.ID == "" {
			spec.ID = spec.Job
		}

		job, err := s.registry.Build(spec.Job)
		if err != nil {
			return err
		}
		if len(spec.Args) > 0 {
			job.SetArgs(spec.Args)
		}

		jobs = append(jobs, scheduledJob{spec: spec, job: job})
		s.logger.Info("job loaded", zap.String("task", spec.ID), zap.String("job", spec.Job))
	}
	s.jobs = jobs
	return nil
}

func (s *Scheduler) LoadAndSchedule(specs []TaskSpec) error {
	if err := s.Load(specs); err != nil {
		return err
	}
	return s.Schedule()
}

func (s *Scheduler) Schedule() error {
	var joined error

	for _, entry := range s.jobs {
		entry := entry

		if entry.spec.Immediate {
			s.logger.Info("running job immediately", zap.String("task", entry.spec.ID), zap.String("job", entry.spec.Job))
			s.run(entry)
		}

		_, err := s.inner.Cron(entry.spec.Cron).Do(func() {
			s.run(entry)
		})
		if err != nil {
			joined = errors.Join(joined, err)
			s.logger.Error("failed to schedule job",
				zap.String("task", entry.spec.ID),
				zap.String("job", entry.spec.Job),
				zap.String("cron", entry.spec.Cron),
				zap.Error(err),
			)
			continue
		}
		s.logger.Info("job scheduled",
			zap.String("task", entry.spec.ID),
			zap.String("job", entry.spec.Job),
			zap.String("cron", entry.spec.Cron),
		)
	}

	return joined
}

func (s *Scheduler) Start() { s.inner.StartAsync() }

func (s *Scheduler) Stop() { s.inner.Stop() }

func (s *Scheduler) run(entry scheduledJob) {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Error("scheduler-level panic",
				zap.String("task", entry.spec.ID),
				zap.String("job", entry.spec.Job),
				zap.Any("panic", r),
			)
		}
	}()

	if err := entry.job.Run(entry.job); err != nil {
		s.logger.Error("job run error",
			zap.String("task", entry.spec.ID),
			zap.String("job", entry.spec.Job),
			zap.Error(err),
		)
	}
}
