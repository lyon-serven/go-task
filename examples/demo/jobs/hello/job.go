package hello

import (
	"fmt"

	task "github.com/lyon-serven/go-task"
	"github.com/lyon-serven/go-task/app"
)

type Job struct {
	*task.BaseJob
}

func NewJob() *Job {
	return &Job{BaseJob: app.NewBaseJob("hello")}
}

func (j *Job) Exec() error {
	greeting, _ := j.Args["greeting"].(string)
	target, _ := j.Args["target"].(string)

	fmt.Printf("%s, %s\n", greeting, target)
	return nil
}
