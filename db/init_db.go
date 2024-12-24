package db

import (
	"context"
	"ecommerce/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"strconv"
	"time"
)

func InitDatabase(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	connString := cfg.ConnString()

	connConfig, err := pgxpool.ParseConfig(connString) // check if the the config is correct
	if err != nil {
		return nil, fmt.Errorf("error parsing db config: %w", err)
	}

	connConfig.MaxConns = int32(cfg.MaxConns)
	connConfig.MaxConnIdleTime = cfg.MaxIdleTime
	connConfig.MaxConnLifetime = cfg.MaxConnLifetime

	pool, err := pgxpool.NewWithConfig(context.Background(), connConfig) // NewWithConfig(instead of new) because of the custom config
	if err != nil {
		return nil, fmt.Errorf("error creating the connection: %w", err)
	}

	// checking the connection
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping db: %w", err)
	}

	log.Println("connected to the database")
	return pool, nil
}

func NewDBConfig() (*config.DBConfig, error) {
	maxIdleTime, err := time.ParseDuration(os.Getenv("DB_MAX_IDLE_TIME"))
	if err != nil {
		return nil, fmt.Errorf("invalid value for DB_MAX_IDLE_TIME: %w", err)
	}

	maxConnLifetime, err := time.ParseDuration(os.Getenv("DB_MAX_CONN_LIFETIME"))
	if err != nil {
		return nil, fmt.Errorf("invalid value for DB_MAX_CONN_LIFETIME: %w", err)
	}

	maxConns, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("error: invalid value for DB_MAX_CONNS %w", err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("error: invalid value for DB_PORT %w", err)
	}

	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" {
		return nil, fmt.Errorf("missing db variables")
	}

	return &config.DBConfig{
		Host:            os.Getenv("DB_HOST"),
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		Name:            os.Getenv("DB_NAME"),
		Port:            port,
		MaxConns:        maxConns,
		MaxIdleTime:     maxIdleTime,
		MaxConnLifetime: maxConnLifetime,
	}, nil
}
