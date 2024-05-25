package redis

import (
	"github.com/go-redis/redis"
	"github.com/pets-shelters/backend-svc/configs"
	"time"
)

type Redis struct {
	client              *redis.Client
	googleStateLifetime time.Duration
	userInfoLifetime    time.Duration
}

func NewRedis(cfg configs.Redis, googleStateLifetime time.Duration, userInfoLifetime time.Duration) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	return &Redis{
		client:              client,
		googleStateLifetime: googleStateLifetime,
		userInfoLifetime:    userInfoLifetime,
	}
}
