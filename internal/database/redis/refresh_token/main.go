package refreshtoken

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
	_redis "github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/redis"
)

type RefreshTokenCache interface {
	Save(ctx context.Context, userId uint) (string, error)
	Remove(ctx context.Context, refreshToken string) error
	GetUserId(ctx context.Context, refreshToken string) (uint, error)
}

type refreshTokenCache struct {
	cfg         config.RefreshTokenConfig
	redisClient *_redis.Client
}

func New(redisClient *_redis.Client) RefreshTokenCache {
	return &refreshTokenCache{
		redisClient: redisClient,
	}
}

func (c *refreshTokenCache) Save(ctx context.Context, userId uint) (string, error) {
	refreshToken := uuid.NewString()
	data := map[string]any{
		"userId":  userId,
		"isValid": true,
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
	value, err := c.redisClient.Get(ctx, refreshToken).Result()
	if err != nil {
		return err
	}
	var data map[string]any
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	data["isValid"] = false
	dataString, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.redisClient.Set(ctx, refreshToken, dataString, c.cfg.ExpirationTime).Err()
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
	isValidVal, ok := data["isValid"]
	if !ok {
		return 0, &redis.InvalidValue{}
	}
	if isValid, ok := isValidVal.(bool); !ok {
		return 0, &redis.IncorrectValueFormat{}
	} else if !isValid {
		return 0, &redis.InvalidValue{}
	}
	userIdVal, ok := data["userId"]
	if !ok {
		return 0, &redis.IncorrectValueFormat{}
	}
	userIdStr, ok := userIdVal.(string)
	if !ok {
		return 0, &redis.IncorrectValueFormat{}
	}
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	return uint(userId), err
}
