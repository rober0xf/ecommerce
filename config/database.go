package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error: loading env file")
	}
}

type DBConfig struct {
	Host            string
	User            string
	Password        string
	Name            string
	Port            int
	MaxConns        int
	MaxIdleTime     time.Duration
	MaxConnLifetime time.Duration
}

func (cfg *DBConfig) ConnString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
