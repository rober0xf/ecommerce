package main

import (
	"ecommerce/config"
	"ecommerce/db"
	"ecommerce/internal/adapters/handler"
	"log"
)

func main() {
	config.LoadEnv()
	dbConfig, err := db.NewDBConfig()
	if err != nil {
		log.Fatalf("error loading db config: %v", err)
	}

	db, err := db.InitDatabase(dbConfig)
	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}
	defer db.Close()

	server := handler.NewAPIServer(":8000", db)
	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
