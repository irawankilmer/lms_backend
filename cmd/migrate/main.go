package main

import (
	"log"

	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db/migrator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Run migrations
	if err := migrator.Migrate(dbInstance); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
