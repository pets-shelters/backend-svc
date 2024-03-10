package redis

import (
	"github.com/go-redis/redis"
	"github.com/pets-shelters/backend-svc/configs"
)

type Redis struct {
	*redis.Client
}

func NewRedis(cfg configs.Redis) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	return &Redis{client}
}
