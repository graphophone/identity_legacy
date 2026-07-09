package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config interface {
	Redis() *RedisConfig
	Postgres() *PostgresConfig
}

type config struct {
	rawConfig *koanf.Koanf
	redis     *RedisConfig
	postgres  *PostgresConfig
}

func LoadConfig(configPath string) (Config, error) {
	rawConfig := koanf.New(".")
	err := rawConfig.Load(file.Provider(configPath), yaml.Parser())

	cfg := config{
		rawConfig: rawConfig,
	}

	cfg.loadRedisConfigFromRaw()
	cfg.loadPostgresConfigFromRaw()

	return &cfg, err
}

func (c *config) Redis() *RedisConfig {
	return c.redis
}

func (c *config) Postgres() *PostgresConfig {
	return c.postgres
}
