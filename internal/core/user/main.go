package user

import (
	"context"

	userdb "graphophone.identity/internal/database/postgres/user"
)

type UserManager interface {
	GetProfile(ctx context.Context, id uint) (*UserProfile, error)
	Register(ctx context.Context, regData *UserRegistrationData) error
	UpdateProfile(ctx context.Context, id uint, userProfile *UserProfile) error
	UpdatePassword(ctx context.Context, id uint, oldPass, newPass string) error
	UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error
	Deactive(ctx context.Context, id uint) error
	Activate(ctx context.Context, id uint) error
}

type userManager struct {
	db userdb.UserDb
}

func New(db userdb.UserDb) UserManager {
	return &userManager{
		db: db,
	}
}

// Activate implements [UserManager].
func (u *userManager) Activate(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// Deactive implements [UserManager].
func (u *userManager) Deactive(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetProfile implements [UserManager].
func (u *userManager) GetProfile(ctx context.Context, id uint) (*UserProfile, error) {
	panic("unimplemented")
}

// Register implements [UserManager].
func (u *userManager) Register(ctx context.Context, regData *UserRegistrationData) error {
	panic("unimplemented")
}

// UpdateAvatar implements [UserManager].
func (u *userManager) UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error {
	panic("unimplemented")
}

// UpdatePassword implements [UserManager].
func (u *userManager) UpdatePassword(ctx context.Context, id uint, oldPass string, newPass string) error {
	panic("unimplemented")
}

// UpdateProfile implements [UserManager].
func (u *userManager) UpdateProfile(ctx context.Context, id uint, userProfile *UserProfile) error {
	panic("unimplemented")
}
