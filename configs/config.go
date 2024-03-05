package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type (
	Config struct {
		HTTP `yaml:"rest"`
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
	filepath, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		log.Fatalf("migrate: environment variable not declared: CONFIG_FILE")
	}

	file, err := os.Open(filepath)
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
