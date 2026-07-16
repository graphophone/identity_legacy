package userdb

import (
	"context"

	"gorm.io/gorm"
	"graphophone.identity/internal/database/postgres"
)

type UserManager interface {
	Get(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	UpdateIsActive(ctx context.Context, id uint, isActive bool) error
	UpdatePasswordHash(ctx context.Context, id uint, passHash string) error
	UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error
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

func (m *userManager) Get(ctx context.Context, id uint) (*User, error) {
	db := m.pg.GetDb()
	user, err := gorm.G[User](db).Where("id = ?", id).First(ctx)
	return &user, err
}

func (m *userManager) Create(ctx context.Context, user *User) (*User, error) {
	db := m.pg.GetDb()
	err := gorm.G[User](db).Create(ctx, user)
	return user, err
}

func (m *userManager) Update(ctx context.Context, user *User) error {
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
		return &postgres.NotFound{}
	}
	return nil
}

func (m *userManager) UpdateIsActive(ctx context.Context, id uint, isActive bool) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "is_active", isActive)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFound{}
	}
	return nil
}

func (m *userManager) UpdatePasswordHash(ctx context.Context, id uint, passHash string) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "password_hash", passHash)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFound{}
	}
	return nil
}

func (m *userManager) UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error {
	db := m.pg.GetDb()
	rows, err := gorm.G[User](db).Where("id = ?", id).Update(ctx, "avatar_url", avatarUrl)
	if err != nil {
		return err
	}
	if rows != 1 {
		return &postgres.NotFound{}
	}
	return nil
}
