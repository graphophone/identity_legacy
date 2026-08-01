package refreshtoken

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/redisdb"
)

type RefreshTokenCache interface {
	Save(ctx context.Context, userId uint) (string, error)
	Remove(ctx context.Context, refreshToken string) error
	GetUserId(ctx context.Context, refreshToken string) (uint, error)
}

type refreshTokenCache struct {
	cfg         config.RefreshTokenConfig
	redisClient *redis.Client
}

func New(redisClient *redis.Client) RefreshTokenCache {
	return &refreshTokenCache{
		redisClient: redisClient,
	}
}

func (c *refreshTokenCache) Save(ctx context.Context, userId uint) (string, error) {
	refreshToken := uuid.NewString()
	data := map[string]any{
		"userId": userId,
	}
	dataString, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	if err := c.redisClient.Set(ctx, refreshToken, dataString, c.cfg.ExpirationTime).Err(); err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (c *refreshTokenCache) Remove(ctx context.Context, refreshToken string) error {
	return c.redisClient.Expire(ctx, refreshToken, time.Millisecond).Err()
}

func (c *refreshTokenCache) GetUserId(ctx context.Context, refreshToken string) (uint, error) {
	value, err := c.redisClient.Get(ctx, refreshToken).Result()
	if err != nil {
		return 0, err
	}
	var data map[string]any
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return 0, err
	}
	userIdVal, ok := data["userId"]
	if !ok {
		return 0, &redisdb.IncorrectValueFormat{}
	}
	userIdFloat, ok := userIdVal.(float64)
	if !ok {
		return 0, &redisdb.IncorrectValueFormat{}
	}
	return uint(userIdFloat), err
}
