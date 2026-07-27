package services

import (
	"context"

	"graphophone.identity/internal/services/common/user"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func NewUserServer() user.UserServiceServer {
	return &UserServer{}
}

func (u *UserServer) ActivateUserProfile(context.Context, *user.ActivateUserProfileRequest) (*user.Empty, error) {
	panic("unimplemented")
}

func (u *UserServer) DeactivateUser(context.Context, *user.DeactivateUserProfileRequest) (*user.Empty, error) {
	panic("unimplemented")
}

func (u *UserServer) GetUserProfile(context.Context, *user.GetProfileRequest) (*user.UserProfile, error) {
	panic("unimplemented")
}

func (u *UserServer) RegisterUser(context.Context, *user.RegisterUserRequest) (*user.UserProfile, error) {
	panic("unimplemented")
}

func (u *UserServer) UpdateUserAvatar(context.Context, *user.UpdateUserAvatarRequest) (*user.Empty, error) {
	panic("unimplemented")
}

func (u *UserServer) UpdateUserPassword(context.Context, *user.UpdateUserPasswordRequest) (*user.Empty, error) {
	panic("unimplemented")
}

func (u *UserServer) UpdateUserProfile(context.Context, *user.UpdateUserProfileRequest) (*user.Empty, error) {
	panic("unimplemented")
}

func (u *UserServer) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}
