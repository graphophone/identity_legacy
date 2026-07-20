package config

import "time"

const (
	jwtKey            = "jwt.key"
	jwtExpirationTime = "jwt.expirationTime"
)

type JwtConfig struct {
	Key            string
	ExpirationTime time.Duration
}

func (c *config) loadJwtConfigFromRaw() {
	key := c.rawConfig.MustString(jwtKey)
	expirationTime := c.rawConfig.MustDuration(jwtExpirationTime)

	c.jwt = &JwtConfig{
		Key:            key,
		ExpirationTime: expirationTime,
	}
}
