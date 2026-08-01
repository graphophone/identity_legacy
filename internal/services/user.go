package services

import (
	"context"

	ccontext "graphophone.identity/internal/context"
	usercore "graphophone.identity/internal/core/user"
	"graphophone.identity/internal/services/common/user"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func NewUserServer() user.UserServiceServer {
	return &UserServer{}
}

func (u *UserServer) ActivateUserProfile(ctx context.Context, req *user.ActivateUserProfileRequest) (*user.Empty, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	if err := um.Activate(ctx, uint(req.UserId)); err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}

func (u *UserServer) DeactivateUser(ctx context.Context, req *user.DeactivateUserProfileRequest) (*user.Empty, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	if err := um.Deactive(ctx, uint(req.UserId)); err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}

func (u *UserServer) GetUserProfile(ctx context.Context, req *user.GetProfileRequest) (*user.UserProfile, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	profile, err := um.GetProfile(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return mapToProfileResponse(profile), nil
}

func (u *UserServer) UpdateUserAvatar(ctx context.Context, req *user.UpdateUserAvatarRequest) (*user.Empty, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	if err := um.UpdateAvatar(ctx, uint(req.UserId), "replace with storage url"); err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}

func (u *UserServer) UpdateUserPassword(ctx context.Context, req *user.UpdateUserPasswordRequest) (*user.Empty, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	if err := um.UpdatePassword(ctx, uint(req.UserId), req.OldPassword, req.NewPassword); err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}

func (u *UserServer) UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*user.Empty, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	updData := usercore.UpdateProfileData{
		Id:        uint(req.UserId),
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Bio:       req.Bio,
		Country:   req.Country,
		City:      req.City,
	}
	if err := um.UpdateProfile(ctx, &updData); err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}

func mapToProfileResponse(profile *usercore.UserProfile) *user.UserProfile {
	return &user.UserProfile{
		Id:        uint32(profile.Id),
		Username:  profile.Username,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Bio:       profile.Bio,
		Country:   profile.Country,
		City:      profile.City,
		AvatarUrl: profile.AvatarUrl,
	}
}
