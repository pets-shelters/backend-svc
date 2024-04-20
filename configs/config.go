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
		Mailjet        `yaml:"mailjet" validate:"required"`
		Twilio         `yaml:"twilio" validate:"required"`
		TemporaryFiles `yaml:"temporary_files" validate:"required"`
		SchedulerHours `yaml:"scheduler_hours" validate:"required"`
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
		WebRedirect   string        `yaml:"web_redirect" validate:"required"`
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
		WriteEndpoint    string `yaml:"write_endpoint" validate:"required"`
		ReadEndpoint     string `yaml:"read_endpoint" validate:"required"`
		PublicReadBucket string `yaml:"public_read_bucket" validate:"required"`
	}

	Mailjet struct {
		PublicKey            string `yaml:"public_key" validate:"required"`
		PrivateKey           string `yaml:"private_key" validate:"required"`
		SenderEmail          string `yaml:"sender_email" validate:"required"`
		SenderName           string `yaml:"sender_name" validate:"required"`
		InvitationTemplateId int    `yaml:"invitation_template_id" validate:"required"`
		InvitationUrl        string `yaml:"invitation_url" validate:"required"`
		TasksTemplateId      int    `yaml:"tasks_template_id" validate:"required"`
		TasksUrl             string `yaml:"tasks_url" validate:"required"`
	}

	Twilio struct {
		Enabled      bool   `yaml:"enabled"`
		AccountSid   string `yaml:"account_sid" validate:"required"`
		AuthToken    string `yaml:"auth_token" validate:"required"`
		SenderNumber string `yaml:"sender_number" validate:"required"`
	}

	TemporaryFiles struct {
		SchedulerPeriod time.Duration `yaml:"scheduler_period" validate:"required"`
		Lifetime        time.Duration `yaml:"lifetime" validate:"required"`
	}

	SchedulerHours struct {
		TasksHour    int64  `yaml:"tasks_hour" validate:"required"`
		WalkingsHour int64  `yaml:"walkings_hour" validate:"required"`
		Location     string `yaml:"location" validate:"required"`
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
