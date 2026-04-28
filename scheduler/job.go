package scheduler

type Job interface {
	Name() string
	Run(job Job) error
	Exec() error
	SetArgs(args map[string]any)
}

type Factory func() Job

type TaskSpec struct {
	ID        string         `yaml:"id" json:"id"`
	Job       string         `yaml:"job" json:"job"`
	Cron      string         `yaml:"cron" json:"cron"`
	Immediate bool           `yaml:"immediate" json:"immediate"`
	Args      map[string]any `yaml:"args" json:"args"`
}
