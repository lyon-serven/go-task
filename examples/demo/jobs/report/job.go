package report

import (
	"fmt"

	"github.com/lyon-serven/go-task/app"
	task "github.com/lyon-serven/go-task"
)

type Job struct {
	*task.BaseJob
}

func NewJob() *Job {
	return &Job{BaseJob: app.NewBaseJob("report")}
}

func (j *Job) Exec() error {
	name, _ := j.Args["name"].(string)
	window, _ := j.Args["window"].(string)
	fmt.Printf("report=%s window=%s\n", name, window)
	return nil
}
