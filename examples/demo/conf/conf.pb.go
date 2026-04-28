package conf

import structpb "google.golang.org/protobuf/types/known/structpb"

type Bootstrap struct {
	App   *App   `json:"app,omitempty"`
	Log   *Log   `json:"log,omitempty"`
	Tasks []*Task `json:"tasks,omitempty"`
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type App struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type Log struct {
	Level        string `json:"level,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
	RotationTime string `json:"rotation_time,omitempty"`
}

func (x *Log) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *Log) GetRotationTime() string {
	if x != nil {
		return x.RotationTime
	}
	return ""
}

type Task struct {
	Id        string           `json:"id,omitempty"`
	Job       string           `json:"job,omitempty"`
	Cron      string           `json:"cron,omitempty"`
	Immediate bool             `json:"immediate,omitempty"`
	Args      *structpb.Struct `json:"args,omitempty"`
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Task) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *Task) GetImmediate() bool {
	if x != nil {
		return x.Immediate
	}
	return false
}

func (x *Task) GetArgs() *structpb.Struct {
	if x != nil {
		return x.Args
	}
	return nil
}
