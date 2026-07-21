package refreshtoken

import _redis "github.com/redis/go-redis/v9"

type RefreshTokenCache interface {
	Save(userId uint, refreshToken string) error
	Remove(refreshToken string) error
	GetUserId(refreshToken string) (uint, error)
}

type refreshTokenCache struct {
	redisClient *_redis.Client
}

func New(redisClient *_redis.Client) RefreshTokenCache {
	return &refreshTokenCache{
		redisClient: redisClient,
	}
}

func (c *refreshTokenCache) Save(userId uint, refreshToken string) error {
	return nil
}

func (c *refreshTokenCache) Remove(refreshToken string) error {
	return nil
}

func (c *refreshTokenCache) GetUserId(refreshToken string) (uint, error) {
	return 0, nil
}
