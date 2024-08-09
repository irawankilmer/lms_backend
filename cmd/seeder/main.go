package main

import (
	"log"
	"time"

	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/models"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance := db.InitDB(cfg)

	// Seed a new user
	now := time.Now()
	user := models.User{
		Username:      "admin",
		Email:         "admin@example.com",
		EmailVerified: &now,
		LastLogin:     &now,
		LastActivity:  &now,
	}

	// Hash password
	if err := user.HashPassword("password123"); err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	if err := dbInstance.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	log.Println("User seeded successfully")
}
