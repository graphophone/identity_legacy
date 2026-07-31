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

func (a *AuthServer) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.Tokens, error) {
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

func (a *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.Tokens, error) {
	am, err := ccontext.GetAuthManager(ctx)
	if err != nil {
		return nil, err
	}
	tokens, err := am.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return mapToGrpcTokens(tokens), nil
}

func (a *AuthServer) Logout(ctx context.Context, tokens *auth.Tokens) (*auth.Empty, error) {
	am, err := ccontext.GetAuthManager(ctx)
	if err != nil {
		return nil, err
	}
	err = am.Logout(ctx, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &auth.Empty{}, nil
}

func (a *AuthServer) LogoutEverywhere(ctx context.Context, tokens *auth.Tokens) (*auth.Empty, error) {
	panic("unimplemented")
}

func (a *AuthServer) RefreshTokens(ctx context.Context, tokens *auth.Tokens) (*auth.Tokens, error) {
	am, err := ccontext.GetAuthManager(ctx)
	if err != nil {
		return nil, err
	}
	newTokens, err := am.Refresh(ctx, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}
	return mapToGrpcTokens(newTokens), nil
}

func (a *AuthServer) ExtractClaims(ctx context.Context, tokens *auth.Tokens) (*auth.Claims, error) {
	am, err := ccontext.GetAuthManager(ctx)
	if err != nil {
		return nil, err
	}
	userId, err := am.ExtractUserId(tokens.AccessToken)
	if err != nil {
		return nil, err
	}
	return &auth.Claims{
		UserId: uint32(userId),
	}, nil
}

func mapToGrpcTokens(tokens *authcore.Tokens) *auth.Tokens {
	return &auth.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
