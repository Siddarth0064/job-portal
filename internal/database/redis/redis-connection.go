package redis

import (
	"fmt"
	"job-portal-api/config"
	// "job-portal-api/internal/services"

	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	cfg := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(":%s", cfg.RedisConfig.Address),
		Password: fmt.Sprintf(":%s", cfg.RedisConfig.RedisPassword),
		DB:       cfg.RedisConfig.Db,
	})

	return rdb
}
