package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	task "github.com/lyon-serven/go-task"
	"github.com/lyon-serven/go-task/conf"
	"github.com/lyon-serven/go-task/internal/bootstrap"
	"gopkg.in/yaml.v3"
)

type demoJob struct {
	*task.BaseJob
}

func (j *demoJob) Exec() error {
	greeting, _ := j.Args["greeting"].(string)
	target, _ := j.Args["target"].(string)
	fmt.Printf("%s, %s\n", greeting, target)
	return nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := bootstrap.NewLogger(&cfg.Log, "")
	if err != nil {
		panic(err)
	}

	registry := task.NewRegistry()
	registry.MustRegister("demo", func() task.Job {
		return &demoJob{BaseJob: task.NewBaseJob("demo", logger)}
	})

	runner := task.New(logger, registry)
	if err := runner.LoadAndSchedule(cfg.Tasks); err != nil {
		panic(err)
	}

	runner.Start()
	defer runner.Stop()

	time.Sleep(2 * time.Second)
}

func loadConfig() (*conf.Config, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("resolve example directory")
	}

	data, err := os.ReadFile(filepath.Join(filepath.Dir(currentFile), "config.yaml"))
	if err != nil {
		return nil, err
	}

	cfg := &conf.Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
