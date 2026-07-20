package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/core"
	"graphophone.identity/internal/core/jwt"
	userdb "graphophone.identity/internal/database/postgres/user"
)

type AuthManager interface {
	Login(ctx context.Context, username, password string) (*Tokens, error)
	Refresh(ctx context.Context, refreshToken string) (*Tokens, error)
	Logout(ctx context.Context, tokens *Tokens) error
}

type authManager struct {
	userDb             userdb.UserDb
	redisClient        *redis.Client
	jwtConfig          *config.JwtConfig
	refreshTokenConfig *config.RefreshTokenConfig
}

func New(userDb userdb.UserDb, redisClient *redis.Client, cfg config.Config) AuthManager {
	return &authManager{
		userDb:             userDb,
		redisClient:        redisClient,
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
	accessToken, err := jwt.GenerateJwtToken(user.ID, user.Username, m.jwtConfig)
	if err != nil {
		return nil, err
	}
	refreshToken := uuid.NewString()
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *authManager) Refresh(ctx context.Context, refreshToken string) (*Tokens, error) {
	return nil, nil
}

func (m *authManager) Logout(ctx context.Context, tokens *Tokens) error {
	return nil
}
