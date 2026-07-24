package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_jwt "github.com/golang-jwt/jwt/v5"
	"graphophone.identity/internal/config"
)

func GenerateJwtToken(id uint, cfg *config.JwtConfig) (string, error) {
	claims := _jwt.RegisteredClaims{
		ExpiresAt: _jwt.NewNumericDate(time.Now().Add(cfg.ExpirationTime)),
		ID:        fmt.Sprintf("%d", id),
	}
	token := _jwt.NewWithClaims(
		_jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString([]byte(cfg.Key))
}

func IsValidJwt(tokenString string, cfg *config.JwtConfig) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(cfg.Key), nil
	})
	if err != nil {
		return false
	}
	expTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false
	}
	return !time.Now().After(expTime.Time)
}

func GetClaims(tokenString string, cfg *config.JwtConfig) (_jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(cfg.Key), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
