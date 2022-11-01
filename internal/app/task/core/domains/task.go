package domains

type Task struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	AbortOnFail bool              `yaml:"abortOnFail"`
	Args        map[string]string `yaml:"args"`
}
