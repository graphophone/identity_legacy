package postgres

import (
	"gorm.io/gorm"
	"graphophone.identity/internal/database/postgres/user"
)

type pgClient struct {
	orm *gorm.DB
}

func (c *pgClient) RegisterModels() error {
	models := []any{
		&user.User{},
	}

	for _, model := range models {
		if err := c.orm.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

func (c *pgClient) Close() error {
	db, err := c.orm.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
