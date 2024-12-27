package main

import (
	"database/sql"
	"ecommerce/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	//	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	config.LoadEnv()

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// temporary connection for migrations
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		log.Fatalf("error connecting to migrations db: %v", err)
	}
	defer db.Close()

	// driver for migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("error to create migrations driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("error to initialize migrations: %v", err)
	}

	// execute migrations
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("error to run migrations: %v", err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("error to run migrations: %v", err)
		}
	}

	log.Println("migrations successfully")
}
