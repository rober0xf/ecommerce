package main

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/db"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config.LoadEnv()
	dbConfig, err := db.NewDBConfig()
	if err != nil {
		log.Fatalf("error loading db config: %v", err)
	}

	// temporary connection for migrations
	db, err := sql.Open("pgx", dbConfig.ConnString())
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
