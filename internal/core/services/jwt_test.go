package services

import (
	"ecommerce/config"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	secret := config.JwtConfig{
		SecretKey:     []byte("secret"),
		SigningMethod: jwt.SigningMethodHS256,
		Expiration:    time.Hour * 24,
	}

	t.Run("token creating", func(t *testing.T) {
		token, err := CreateJWT(secret, 1)
		if err != nil {
			t.Fatalf("error creating jwt: %v", err)
		}

		if token == "" {
			t.Error("expected token, got nothing")
		}

	})

	t.Run("token validation", func(t *testing.T) {
		token, _ := CreateJWT(secret, 1)

		// parse and validate token
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return secret.SecretKey, nil
		})

		if err != nil {
			t.Fatalf("expected valid token, got error: %v", err)
		}

		if !parsedToken.Valid {
			t.Error("invalid token")
		}
	})
}
