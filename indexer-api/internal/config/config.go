package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct for webapp config
type Config struct {
	Kafka struct {
		Address   string `yaml:"address"`
		Topic     string `yaml:"topic"`
		Partition int    `yaml:"partition"`
	} `yaml:"kafka"`

	ProductAPI struct {
		Address string `yaml:"address"`
		Path    string `yaml:"path"`
	} `yaml:"productAPI"`

	SearchAPI struct {
		Address string `yaml:"address"`
		Path    string `yaml:"path"`
	} `yaml:"searchAPI"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	err := validateConfigPath(configPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}