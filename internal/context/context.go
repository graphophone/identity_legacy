package ccontext

import (
	"context"

	"graphophone.identity/internal/config"
	authcore "graphophone.identity/internal/core/auth"
	usercore "graphophone.identity/internal/core/user"
	"graphophone.identity/internal/database/postgres"
	userdb "graphophone.identity/internal/database/postgres/user"
	"graphophone.identity/internal/database/redis"
	refreshtoken "graphophone.identity/internal/database/redis/refresh_token"
)

type customContext struct {
	context.Context
	cfg         config.Config
	userManager usercore.UserManager
	authManager authcore.AuthManager
}

func newCustomContext(cfg config.Config) (*customContext, error) {
	baseCtx := context.Background()

	pg, err := postgres.New(cfg.Postgres())
	if err != nil {
		return nil, err
	}
	userDb, err := userdb.New(pg)
	if err != nil {
		return nil, err
	}
	userManager := usercore.New(userDb)

	redisClient, err := redis.Connect(baseCtx, cfg.Redis())
	if err != nil {
		return nil, err
	}
	refreshTokenCache := refreshtoken.New(redisClient)
	authManager := authcore.New(userDb, refreshTokenCache, cfg)

	return &customContext{
		cfg:         cfg,
		userManager: userManager,
		authManager: authManager,
	}, nil
}
