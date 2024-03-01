package configs

import (
	"fmt"
	"os"
)

type (
	// Config -.
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"log_level"`
	}

	PG struct {
		URL string `yaml:"db_url"`
	}
)

func NewConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
