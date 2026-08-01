package refreshtoken

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/redisdb"
)

type RefreshTokenCache interface {
	Save(ctx context.Context, userId uint) (string, error)
	Remove(ctx context.Context, refreshToken string) error
	RemoveAll(ctx context.Context, userId uint) error
	GetUserId(ctx context.Context, refreshToken string) (uint, error)
}

type refreshTokenCache struct {
	cfg         config.RefreshTokenConfig
	redisClient *redis.Client
}

func New(ctx context.Context, redisClient *redis.Client) (RefreshTokenCache, error) {
	_, err := redisClient.FTCreate(
		ctx, "idx:refresh_tokens",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []any{"refresh_tokens"},
		},
		&redis.FieldSchema{
			FieldName: "$.userId",
			As:        "userId",
			FieldType: redis.SearchFieldTypeNumeric,
		},
	).Result()
	if err != nil {
		return nil, err
	}
	return &refreshTokenCache{
		redisClient: redisClient,
	}, nil
}

func (c *refreshTokenCache) Save(ctx context.Context, userId uint) (string, error) {
	refreshToken := uuid.NewString()
	data := map[string]any{
		"userId": userId,
	}
	key := getKey(refreshToken)
	if err := c.redisClient.JSONSet(ctx, key, "$", data).Err(); err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (c *refreshTokenCache) Remove(ctx context.Context, refreshToken string) error {
	key := getKey(refreshToken)
	return c.redisClient.Expire(ctx, key, time.Second).Err()
}

func (c *refreshTokenCache) RemoveAll(ctx context.Context, userId uint) error {
	entries, err := c.redisClient.FTSearchWithArgs(
		ctx, "idx:refresh_tokens",
		strconv.FormatUint(uint64(userId), 10),
		&redis.FTSearchOptions{
			Return: []redis.FTSearchReturn{
				{
					FieldName: "$.userId",
					As:        "userId",
				},
			},
		},
	).Result()
	if err != nil {
		return err
	}
	for _, entry := range entries.Docs {
		fmt.Println(entry.ID)
	}
	return nil
}

func (c *refreshTokenCache) GetUserId(ctx context.Context, refreshToken string) (uint, error) {
	key := getKey(refreshToken)
	value, err := c.redisClient.Get(ctx, key).Result()
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
	userIdStr, ok := userIdVal.(string)
	if !ok {
		return 0, &redisdb.IncorrectValueFormat{}
	}
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	return uint(userId), err
}

func getKey(refreshToken string) string {
	return "refresh_token:" + refreshToken
}
