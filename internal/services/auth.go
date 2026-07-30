package services

import (
	"context"

	ccontext "graphophone.identity/internal/context"
	authcore "graphophone.identity/internal/core/auth"
	usercore "graphophone.identity/internal/core/user"
	"graphophone.identity/internal/services/common/auth"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthServer() auth.AuthServiceServer {
	return &AuthServer{}
}

func (u *UserServer) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.Tokens, error) {
	um, err := ccontext.GetUserManager(ctx)
	if err != nil {
		return nil, err
	}
	am, err := ccontext.GetAuthManager(ctx)
	if err != nil {
		return nil, err
	}
	regData := usercore.RegisterUserData{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	_, err = um.Register(ctx, &regData)
	if err != nil {
		return nil, err
	}
	tokens, err := am.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return mapToGrpcTokens(tokens), nil
}

func (a *AuthServer) Login(context.Context, *auth.LoginRequest) (*auth.Tokens, error) {
	panic("unimplemented")
}

func (a *AuthServer) Logout(context.Context, *auth.Tokens) (*auth.Empty, error) {
	panic("unimplemented")
}

func (a *AuthServer) LogoutEverywhere(context.Context, *auth.Tokens) (*auth.Empty, error) {
	panic("unimplemented")
}

func (a *AuthServer) RefreshTokens(context.Context, *auth.Tokens) (*auth.Tokens, error) {
	panic("unimplemented")
}

func (a *AuthServer) SignUp(context.Context, *auth.SignUpRequest) (*auth.Tokens, error) {
	panic("unimplemented")
}

func (a *AuthServer) VerifyTokens(context.Context, *auth.Tokens) (*auth.Empty, error) {
	panic("unimplemented")
}

func mapToGrpcTokens(tokens *authcore.Tokens) *auth.Tokens {
	return &auth.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
