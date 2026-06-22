package auth

import (
	"time"

	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(userID string, email string, cfg *config.Config) (string, error) {
	expiration := time.Now().Add(time.Hour * 24 * time.Duration(cfg.JWTRefreshExpirationDays))

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
