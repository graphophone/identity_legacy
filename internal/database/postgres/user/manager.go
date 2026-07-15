package user

import (
	"context"

	"gorm.io/gorm"
	"graphophone.identity/internal/database/postgres"
)

type UserManager interface {
	Get(ctx context.Context, id uint) (User, error)
	Register(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	Deactivate(ctx context.Context, id uint) error
}

type userManager struct {
	pg postgres.PgClient
}

func New(pg postgres.PgClient) (UserManager, error) {
	if err := pg.RegisterModel(&User{}); err != nil {
		return nil, err
	}
	return &userManager{
		pg: pg,
	}, nil
}

func (m *userManager) Get(ctx context.Context, id uint) (User, error) {
	db := m.pg.GetDb()
	return gorm.G[User](db).Where("id = ?", id).First(ctx)
}

func (m *userManager) Register(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (m *userManager) Update(ctx context.Context, user *User) error {
	return nil
}

func (m *userManager) Deactivate(ctx context.Context, id uint) error {
	return nil
}
