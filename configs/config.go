package configs

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
	"log"
	"os"
)

type (
	Config struct {
		HTTP `yaml:"http" validate:"required"`
		Log  `yaml:"logger" validate:"required"`
		PG   `yaml:"postgres" validate:"required"`
	}

	HTTP struct {
		Addr string `yaml:"addr" validate:"required"`
	}

	Log struct {
		Level string `yaml:"log_level" validate:"required"`
	}

	PG struct {
		URL string `yaml:"db_url" validate:"required"`
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
	validate := validator.New()
	decoder := yaml.NewDecoder(
		file,
		yaml.Validator(validate),
		yaml.Strict(),
	)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
