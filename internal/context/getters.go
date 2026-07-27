package context

import (
	"context"

	"graphophone.identity/internal/core/auth"
	"graphophone.identity/internal/core/user"
)

func GetUserManager(ctx context.Context) (user.UserManager, error) {
	customContext, ok := ctx.(customContext)
	if !ok {
		return nil, &InvalidContextType{}
	}
	return customContext.userManager, nil
}

func GetAuthManager(ctx context.Context) (auth.AuthManager, error) {
	customContext, ok := ctx.(customContext)
	if !ok {
		return nil, &InvalidContextType{}
	}
	return customContext.authManager, nil
}
