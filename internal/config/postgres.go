package config

import (
	"time"
)

const (
	postgresHost     = "postgres.host"
	postgresPort     = "postgres.port"
	postgresUser     = "postgres.user"
	postgresPassword = "postgres.password"
	postgresDatabase = "postgres.database"
	postgresTimeout  = "postgres.timeout"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Timeout  time.Duration
}

func (c *config) loadPostgresConfigFromRaw() {
	host := c.rawConfig.MustString(postgresHost)
	port := c.rawConfig.MustInt(postgresPort)
	user := c.rawConfig.String(postgresUser)
	password := c.rawConfig.MustString(postgresPassword)
	database := c.rawConfig.MustString(postgresDatabase)
	timeout := c.rawConfig.Duration(postgresTimeout)

	c.postgres = &PostgresConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		Timeout:  timeout,
	}
}
