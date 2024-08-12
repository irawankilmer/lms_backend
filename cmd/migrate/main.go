package main

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/models"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Define the migrations
	migrations := []*gormigrate.Migration{
		{
			ID: "20240812_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "20240812_create_reset_tokens_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.ResetToken{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("reset_tokens")
			},
		},
	}

	// Initialize gormigrate
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	// Run the migrations
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Migration completed successfully")
}
