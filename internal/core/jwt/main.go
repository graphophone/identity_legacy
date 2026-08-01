package jwtcore

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"graphophone.identity/internal/config"
)

func GenerateJwtToken(id uint, cfg *config.JwtConfig) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpirationTime)),
		Subject:   fmt.Sprintf("%d", id),
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString([]byte(cfg.Key))
}

func GetClaims(tokenString string, cfg *config.JwtConfig) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(cfg.Key), nil
	})
	if err != nil {
		return nil, &InvalidJwt{}
	}
	expTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, &InvalidJwt{}
	}
	if time.Now().After(expTime.Time) {
		return nil, &ExpiredJwt{}
	}
	return token.Claims, nil
}
