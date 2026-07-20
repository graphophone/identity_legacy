package auth

import (
	"github.com/redis/go-redis/v9"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/core/user"
)

type AuthManager interface {
	Login(username, password string) (*Tokens, error)
	Register(regData *user.RegisterUserData) (*Tokens, error)
	Refresh(refreshToken string) (*Tokens, error)
	Logout(tokens *Tokens) error
}

type authManager struct {
	jwtConfig          *config.JwtConfig
	refreshTokenConfig *config.RefreshTokenConfig
	redisClient        *redis.Client
}

func New(redisClient *redis.Client, cfg config.Config) AuthManager {
	return &authManager{
		jwtConfig:          cfg.Jwt(),
		refreshTokenConfig: cfg.RefreshToken(),
		redisClient:        redisClient,
	}
}

func (m *authManager) Login(username, password string) (*Tokens, error) {
	return nil, nil
}

func (m *authManager) Register(regData *user.RegisterUserData) (*Tokens, error) {
	return nil, nil
}

func (m *authManager) Refresh(refreshToken string) (*Tokens, error) {
	return nil, nil
}

func (m *authManager) Logout(tokens *Tokens) error {
	return nil
}
