package config

import "fmt"

const (
	redisHost     = "redis.host"
	redisPort     = "redis.port"
	redisPassword = "redis.password"
	redisDatabase = "redis.database"
)

type RedisConfig struct {
	Address  string
	Password string
	Database int
}

func (c *config) loadRedisConfigFromRaw() {
	host := c.rawConfig.MustString(redisHost)
	port := c.rawConfig.MustInt(redisPort)
	password := c.rawConfig.MustString(redisPassword)
	database := c.rawConfig.MustInt(redisDatabase)

	c.redis = &RedisConfig{
		Address:  fmt.Sprintf("%s:%d", host, port),
		Password: password,
		Database: database,
	}
}
