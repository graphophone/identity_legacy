package usercore

import (
	"context"

	"gorm.io/gorm"
	"graphophone.identity/internal/core"
	userdb "graphophone.identity/internal/database/postgresdb/user"
)

type UserManager interface {
	GetProfile(ctx context.Context, id uint) (*UserProfile, error)
	Register(ctx context.Context, regData *RegisterUserData) (*UserProfile, error)
	UpdateProfile(ctx context.Context, updData *UpdateProfileData) error
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

func (u *userManager) Activate(ctx context.Context, id uint) error {
	return u.db.UpdateIsActive(ctx, id, true)
}

func (u *userManager) Deactive(ctx context.Context, id uint) error {
	return u.db.UpdateIsActive(ctx, id, false)
}

func (u *userManager) GetProfile(ctx context.Context, id uint) (*UserProfile, error) {
	user, err := u.db.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return userProfileFromUser(user), nil
}

func (u *userManager) Register(ctx context.Context, regData *RegisterUserData) (*UserProfile, error) {
	userModel, err := userFromRegisterUserData(regData)
	if err != nil {
		return nil, err
	}
	user, err := u.db.Create(ctx, userModel)
	if err != nil {
		return nil, err
	}
	return userProfileFromUser(user), nil
}

func (u *userManager) UpdateAvatar(ctx context.Context, id uint, avatarUrl string) error {
	return u.db.UpdateAvatar(ctx, id, avatarUrl)
}

func (u *userManager) UpdatePassword(ctx context.Context, id uint, oldPass string, newPass string) error {
	user, err := u.db.Get(ctx, id)
	if err != nil {
		return err
	}
	if !core.IsPasswordValid(oldPass, user.PasswordHash) {
		return &core.IncorrectPasswordErr{}
	}
	newPassHash, err := core.HashPassword(newPass)
	if err != nil {
		return err
	}
	return u.db.UpdatePasswordHash(ctx, id, newPassHash)
}

func (u *userManager) UpdateProfile(ctx context.Context, updData *UpdateProfileData) error {
	return u.db.Update(ctx, userFromUpdateUserData(updData))
}

func userProfileFromUser(user *userdb.User) *UserProfile {
	return &UserProfile{
		Id:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Bio:       user.Bio,
		Country:   user.Country,
		City:      user.City,
		AvatarUrl: user.AvatarUrl,
	}
}

func userFromUpdateUserData(updData *UpdateProfileData) *userdb.User {
	return &userdb.User{
		Model: gorm.Model{
			ID: updData.Id,
		},
		Username:  updData.Username,
		FirstName: updData.FirstName,
		LastName:  updData.LastName,
		Bio:       updData.Bio,
		Country:   updData.Country,
		City:      updData.City,
	}
}

func userFromRegisterUserData(regData *RegisterUserData) (*userdb.User, error) {
	passHash, err := core.HashPassword(regData.Password)
	if err != nil {
		return nil, err
	}
	return &userdb.User{
		Username:     regData.Username,
		Email:        regData.Email,
		PasswordHash: passHash,
		FirstName:    regData.FirstName,
		LastName:     regData.LastName,
	}, nil
}
