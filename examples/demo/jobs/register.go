package jobs

import (
	"github.com/lyon-serven/go-task/examples/demo/jobs/hello"
	"github.com/lyon-serven/go-task/examples/demo/jobs/report"
	"github.com/lyon-serven/go-task/scheduler"
)

func init() {
	scheduler.Register(
		hello.NewJob,
		report.NewJob,
	)
}
