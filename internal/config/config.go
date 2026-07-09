package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config interface {
	Redis() *RedisConfig
}

type config struct {
	rawConfig *koanf.Koanf
	redis     *RedisConfig
}

func (c *config) Redis() *RedisConfig {
	return c.redis
}

func LoadConfig(configPath string) (Config, error) {
	rawConfig := koanf.New(".")
	err := rawConfig.Load(file.Provider(configPath), yaml.Parser())

	cfg := config{
		rawConfig: rawConfig,
	}

	cfg.loadRedisConfigFromRaw()

	return &cfg, err
}
