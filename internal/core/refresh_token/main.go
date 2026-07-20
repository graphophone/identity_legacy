package refreshtoken

import (
	"github.com/google/uuid"
	"graphophone.identity/internal/config"
)

func GenerateRefreshToken(cfg config.JwtConfig) string {
	return uuid.NewString()
}
