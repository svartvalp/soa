package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for webapp config
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	DatabaseDsn string `yaml:"databaseDsn"`

	S3 struct {
		Bucket          string `yaml:"bucket"`
		URL             string `yaml:"url"`
		Region          string `yaml:"region"`
		AccessKeyID     string `yaml:"accessKeyID"`
		SecretAccessKey string `yaml:"secretAccessKey"`
	} `yaml:"s3"`

	Kafka struct {
		Address   string `yaml:"address"`
		Topic     string `yaml:"topic"`
		Partition int    `yaml:"partition"`
	} `yaml:"kafka"`
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
