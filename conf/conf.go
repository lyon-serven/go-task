package conf

import task "github.com/lyon-serven/go-task"

type Config struct {
	App   App             `yaml:"app"`
	Log   Log             `yaml:"log"`
	Tasks []task.TaskSpec `yaml:"tasks"`
}

type App struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

type Log struct {
	Level        string `yaml:"level"`
	FilePath     string `yaml:"file_path"`
	FileName     string `yaml:"file_name"`
	FilePattern  string `yaml:"file_pattern"`
	RotationTime string `yaml:"rotation_time"`
	FileAge      string `yaml:"file_age"`
	Environment  string `yaml:"environment"`
	SystemName   string `yaml:"system_name"`
}
