package authcore

import (
	"context"
	"strconv"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/core"
	jwtcore "graphophone.identity/internal/core/jwt"
	userdb "graphophone.identity/internal/database/postgres/user"
	refreshtoken "graphophone.identity/internal/database/redis/refresh_token"
)

type AuthManager interface {
	Login(ctx context.Context, username, password string) (*Tokens, error)
	Refresh(ctx context.Context, refreshToken string) (*Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	ExtractUserId(accessToken string) (uint, error)
}

type authManager struct {
	userDb             userdb.UserDb
	refreshTokenCache  refreshtoken.RefreshTokenCache
	jwtConfig          *config.JwtConfig
	refreshTokenConfig *config.RefreshTokenConfig
}

func New(
	userDb userdb.UserDb,
	refreshTokenCache refreshtoken.RefreshTokenCache,
	cfg config.Config,
) AuthManager {
	return &authManager{
		userDb:             userDb,
		refreshTokenCache:  refreshTokenCache,
		jwtConfig:          cfg.Jwt(),
		refreshTokenConfig: cfg.RefreshToken(),
	}
}

func (m *authManager) Login(ctx context.Context, username, password string) (*Tokens, error) {
	user, err := m.userDb.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !core.IsPasswordValid(password, user.PasswordHash) {
		return nil, &core.IncorrectPasswordErr{}
	}
	return m.generateTokens(ctx, user.ID)
}

func (m *authManager) Refresh(ctx context.Context, refreshToken string) (*Tokens, error) {
	userId, err := m.refreshTokenCache.GetUserId(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if err := m.refreshTokenCache.Remove(ctx, refreshToken); err != nil {
		return nil, err
	}
	return m.generateTokens(ctx, userId)
}

func (m *authManager) Logout(ctx context.Context, refreshToken string) error {
	return m.refreshTokenCache.Remove(ctx, refreshToken)
}

func (m *authManager) ExtractUserId(accessToken string) (uint, error) {
	claims, err := jwtcore.GetClaims(accessToken, m.jwtConfig)
	if err != nil {
		return 0, err
	}
	userIdStr, err := claims.GetSubject()
	if err != nil {
		return 0, err
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, err
	}
	return uint(userId), nil
}

func (m *authManager) generateTokens(ctx context.Context, userId uint) (*Tokens, error) {
	accessToken, err := jwtcore.GenerateJwtToken(userId, m.jwtConfig)
	if err != nil {
		return nil, err
	}
	refreshToken, err := m.refreshTokenCache.Save(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
