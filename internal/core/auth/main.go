package auth

import "graphophone.identity/internal/core/user"

type AuthManager interface {
	Login(username, password string) (Tokens, error)
	Register(regData user.RegisterUserData) (Tokens, error)
	Refresh(refreshToken string) (Tokens, error)
	Logout(tokens Tokens) error
}
