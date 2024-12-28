package services

import (
	"ecommerce/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(cfg config.JwtConfig, userID int) (string, error) {
	token := jwt.NewWithClaims(cfg.SigningMethod, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(cfg.Expiration).Unix(),
	})

	// sign the token with the secret key
	tokenString, err := token.SignedString(cfg.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
