package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config interface {
	Redis() *RedisConfig
	Postgres() *PostgresConfig
	Jwt() *JwtConfig
	RefreshToken() *RefreshTokenConfig
}

type config struct {
	rawConfig    *koanf.Koanf
	redis        *RedisConfig
	postgres     *PostgresConfig
	jwt          *JwtConfig
	refreshToken *RefreshTokenConfig
}

func LoadConfig(configPath string) (Config, error) {
	rawConfig := koanf.New(".")
	err := rawConfig.Load(file.Provider(configPath), yaml.Parser())

	cfg := config{
		rawConfig: rawConfig,
	}

	cfg.loadRedisConfigFromRaw()
	cfg.loadPostgresConfigFromRaw()
	cfg.loadJwtConfigFromRaw()
	cfg.loadRefreshTokenConfigFromRaw()

	return &cfg, err
}

func (c *config) Redis() *RedisConfig {
	return c.redis
}

func (c *config) Postgres() *PostgresConfig {
	return c.postgres
}

func (c *config) Jwt() *JwtConfig {
	return c.jwt
}

func (c *config) RefreshToken() *RefreshTokenConfig {
	return c.refreshToken
}
