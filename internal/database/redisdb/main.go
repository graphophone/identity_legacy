package redisdb

import (
	"context"

	_redis "github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
)

func Connect(ctx context.Context, cfg *config.RedisConfig) (*_redis.Client, error) {
	client := _redis.NewClient(&_redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	if err := client.Set(ctx, "testKey", "testVal", 0).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
