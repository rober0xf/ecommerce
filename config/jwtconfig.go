package config

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type JwtConfig struct {
	SecretKey     []byte
	SigningMethod jwt.SigningMethod
	Expiration    time.Duration
}

func LoadJWTConfig() *JwtConfig {
	LoadEnv()

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("must have a jwt key in environment file")
	}

	expirationStr := os.Getenv("JWT_EXPIRATION")
	if expirationStr == "" {
		expirationStr = "168h"
	}

	expiration, err := time.ParseDuration(expirationStr)
	if err != nil {
		log.Fatalf("invalid jwt expiration format: %v", err)
	}

	return &JwtConfig{
		SecretKey:     []byte(secretKey),
		SigningMethod: jwt.SigningMethodHS256,
		Expiration:    expiration,
	}
}
