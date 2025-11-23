package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/example/graphql-mysql-api/pkg/config"
	"github.com/example/graphql-mysql-api/pkg/database"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	log.Println("database migration completed successfully")
}
