package config

import "time"

type RefreshTokenConfig struct {
	ExpirationTime time.Duration
}

const refreshTokenExpirationTime = "refreshToken.expirationTime"

func (c *config) loadRefreshTokenConfigFromRaw() {
	expirationTime := c.rawConfig.MustDuration(refreshTokenExpirationTime)

	c.refreshToken = &RefreshTokenConfig{
		ExpirationTime: expirationTime,
	}
}
