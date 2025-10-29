package redis

import (
	"userservice/configs"

	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg *configs.Configs) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
}
