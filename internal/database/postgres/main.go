package postgres

import (
	"database/sql"

	"github.com/lib/pq"
	"graphophone.identity/internal/config"
)

func Connect(cfg *config.PostgresConfig) (*sql.DB, error) {
	postgresCfg := pq.Config{
		Host:           cfg.Host,
		Port:           uint16(cfg.Port),
		User:           cfg.User,
		Password:       cfg.Password,
		Database:       cfg.Database,
		ConnectTimeout: cfg.Timeout,
		SSLMode:        pq.SSLModeDisable,
	}

	connector, err := pq.NewConnectorConfig(postgresCfg)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)
	return db, db.Ping()
}
