package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
)

func Connect(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	if err := client.Set(ctx, "testKey", "testVal", 0).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
