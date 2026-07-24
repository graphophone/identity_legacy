package userdb

import (
	"context"

	"gorm.io/gorm"
	"graphophone.identity/internal/database/postgres"
)

type UserDb interface {
	Get(ctx context.Context, id uint) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	UpdateIsActive(ctx context.Context, id uint, isActive bool) error
	UpdatePasswordHash(ctx context.Context, id uint, passHash string) error
	UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error
}

type userDb struct {
	pg postgres.PgClient
}

func New(pg postgres.PgClient) (UserDb, error) {
	if err := pg.RegisterModel(&User{}); err != nil {
		return nil, err
	}
	return &userDb{
		pg: pg,
	}, nil
}

func (m *userDb) Get(ctx context.Context, id uint) (*User, error) {
	db := m.pg.GetDb()
	user, err := gorm.G[User](db).Where("id = ?", id).First(ctx)
	return &user, err
}

func (m *userDb) GetByUsername(ctx context.Context, username string) (*User, error) {
	db := m.pg.GetDb()
	user, err := gorm.G[User](db).Where("username = ?", username).First(ctx)
	return &user, err
}

func (m *userDb) Create(ctx context.Context, user *User) (*User, error) {
	db := m.pg.GetDb()
	err := gorm.G[User](db).Create(ctx, user)
	return user, err
}

func (m *userDb) Update(ctx context.Context, user *User) error {
	newFields := map[string]any{
		"username":   user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"bio":        user.Bio,
		"country":    user.Country,
		"city":       user.City,
	}
	db := m.pg.GetDb()
	rows, err := gorm.G[map[string]any](db).
		Table("users").
		Where("id = ?", user.ID).
		Updates(ctx, newFields)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFoundErr{}
	}
	return nil
}

func (m *userDb) UpdateIsActive(ctx context.Context, id uint, isActive bool) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "is_active", isActive)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFoundErr{}
	}
	return nil
}

func (m *userDb) UpdatePasswordHash(ctx context.Context, id uint, passHash string) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "password_hash", passHash)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFoundErr{}
	}
	return nil
}

func (m *userDb) UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "avatar_url", avatarUrl)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFoundErr{}
	}
	return nil
}
