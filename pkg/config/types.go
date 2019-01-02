package config

// Config represents the cli config
type Config struct {
	Version  string
	TaskFile string
}

func newConfig() *Config {
	return &Config{
		Version:  "1",
		TaskFile: "tasks",
	}
}
