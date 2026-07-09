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
	host := c.rawConfig.String(redisHost)
	port := c.rawConfig.Int(redisPort)
	password := c.rawConfig.String(redisPassword)
	database := c.rawConfig.Int(redisDatabase)

	c.redis = &RedisConfig{
		Address:  fmt.Sprintf("%s:%d", host, port),
		Password: password,
		Database: database,
	}
}
