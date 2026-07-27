package context

import (
	"context"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/core/auth"
	"graphophone.identity/internal/core/user"
	"graphophone.identity/internal/database/postgres"
	userdb "graphophone.identity/internal/database/postgres/user"
	"graphophone.identity/internal/database/redis"
	refreshtoken "graphophone.identity/internal/database/redis/refresh_token"
)

type Context interface {
	context.Context
	GetUserManager() user.UserManager
	GetAuthManager() auth.AuthManager
}

type customContext struct {
	context.Context
	cfg         config.Config
	userManager user.UserManager
	authManager auth.AuthManager
}

func NewContext(cfg config.Config) (Context, error) {
	baseCtx := context.Background()

	pg, err := postgres.New(cfg.Postgres())
	if err != nil {
		return nil, err
	}
	userDb, err := userdb.New(pg)
	if err != nil {
		return nil, err
	}
	userManager := user.New(userDb)

	redisClient, err := redis.Connect(baseCtx, cfg.Redis())
	if err != nil {
		return nil, err
	}
	refreshTokenCache := refreshtoken.New(redisClient)
	authManager := auth.New(userDb, refreshTokenCache, cfg)

	return &customContext{
		cfg:         cfg,
		Context:     baseCtx,
		userManager: userManager,
		authManager: authManager,
	}, nil
}

func (c *customContext) GetAuthManager() auth.AuthManager {
	return c.GetAuthManager()
}

func (c *customContext) GetUserManager() user.UserManager {
	return c.GetUserManager()
}
