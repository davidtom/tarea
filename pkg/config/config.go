package config

import "path/filepath"

// TODO: add methods/cmd to make config configurable
// TODO: create an 'init' cmd to initialize the config file
// TODO: model this off go-t library, BUT DON'T COPY - use only what I actually need (I can add more later, PLUS, who knows how good that code is -> make my own choices)

const (
	configDir  = ".config/"
	configFile = "config.json"
)

func getConfigPath() (string, error) {
	return filepath.Join(configDir, configFile), nil
}

// LoadConfig returns the saved config file
func LoadConfig() *Config {
	// TODO: build this out
	return newConfig()
}
