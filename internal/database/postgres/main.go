package postgres

import (
	"database/sql"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"graphophone.identity/internal/config"
)

type PgClient interface {
	RegisterModels() error
	Close() error
}

func New(cfg *config.PostgresConfig) (PgClient, error) {
	pgCfg := pq.Config{
		Host:           cfg.Host,
		Port:           uint16(cfg.Port),
		User:           cfg.User,
		Password:       cfg.Password,
		Database:       cfg.Database,
		ConnectTimeout: cfg.Timeout,
		SSLMode:        pq.SSLModeDisable,
	}

	connector, err := pq.NewConnectorConfig(pgCfg)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &pgClient{
		orm: orm,
	}, nil
}
