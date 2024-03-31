package configs

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
	"log"
	"os"
	"time"
)

type (
	Config struct {
		HTTP           `yaml:"http" validate:"required"`
		Log            `yaml:"logger" validate:"required"`
		PG             `yaml:"postgres" validate:"required"`
		OAuth          `yaml:"oauth" validate:"required"`
		Jwt            `yaml:"jwt" validate:"required"`
		Infrastructure `yaml:"infrastructure" validate:"required"`
		Redis          `yaml:"redis" validate:"required"`
		S3             `yaml:"s3" validate:"required"`
		TemporaryFiles `yaml:"temporary_files" validate:"required"`
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

	OAuth struct {
		ClientID      string        `yaml:"client_id" validate:"required"`
		ClientSecret  string        `yaml:"client_secret" validate:"required"`
		StateLifetime time.Duration `yaml:"state_lifetime" validate:"required"`
	}

	Jwt struct {
		AccessSecret    string        `yaml:"access_secret" validate:"required"`
		AccessLifetime  time.Duration `yaml:"access_lifetime" validate:"required"`
		RefreshSecret   string        `yaml:"refresh_secret" validate:"required"`
		RefreshLifetime time.Duration `yaml:"refresh_lifetime" validate:"required"`
	}

	Infrastructure struct {
		ServiceUrl   string `yaml:"service_url" validate:"required"`
		WebClientUrl string `yaml:"webclient_url" validate:"required"`
		Domain       string `yaml:"domain" validate:"required"`
	}

	Redis struct {
		Addr     string `yaml:"addr" validate:"required"`
		Password string `yaml:"password" validate:"required"`
	}

	S3 struct {
		Region           string `yaml:"region" validate:"required"`
		AccessKey        string `yaml:"access_key" validate:"required"`
		SecretKey        string `yaml:"secret_key" validate:"required"`
		Endpoint         string `yaml:"endpoint" validate:"-"`
		PublicReadBucket string `yaml:"public_read_bucket" validate:"required"`
	}

	TemporaryFiles struct {
		SchedulerPeriod time.Duration `yaml:"scheduler_period" validate:"required"`
		Lifetime        time.Duration `yaml:"lifetime" validate:"required"`
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
