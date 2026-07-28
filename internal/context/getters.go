package ccontext

import (
	"context"

	authcore "graphophone.identity/internal/core/auth"
	usercore "graphophone.identity/internal/core/user"
)

func GetUserManager(ctx context.Context) (usercore.UserManager, error) {
	customContext, ok := ctx.(customContext)
	if !ok {
		return nil, &InvalidContextType{}
	}
	return customContext.userManager, nil
}

func GetAuthManager(ctx context.Context) (authcore.AuthManager, error) {
	customContext, ok := ctx.(customContext)
	if !ok {
		return nil, &InvalidContextType{}
	}
	return customContext.authManager, nil
}
