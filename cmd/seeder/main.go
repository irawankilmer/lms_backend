package main

import (
	"log"

	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/db/seeder"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance := db.InitDB(cfg)

	// Run the centralized seed function
	if err := seeder.Seed(dbInstance); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}
}
