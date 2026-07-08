package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func LoadConfig(configPath string) (*koanf.Koanf, error) {
	config := koanf.New("/")
	err := config.Load(file.Provider(configPath), yaml.Parser())
	return config, err
}
