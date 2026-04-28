package main

import (
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	task "github.com/lyon-serven/go-task"
	"github.com/lyon-serven/go-task/app"
	_ "github.com/lyon-serven/go-task/examples/demo/jobs"
	"github.com/lyon-serven/go-task/scheduler"
)

func main() {
	app.Init(defaultConfigPath())

	registry := scheduler.NewRegisteredRegistry()
	runner := task.New(app.Logger, registry)
	if err := runner.LoadAndSchedule(app.Config.Tasks); err != nil {
		panic(err)
	}

	runner.Start()
	defer runner.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
	case <-time.After(3 * time.Second):
	}
}

func defaultConfigPath() string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "conf/config.yaml"
	}
	return filepath.Join(filepath.Dir(currentFile), "..", "..", "conf", "config.yaml")
}
