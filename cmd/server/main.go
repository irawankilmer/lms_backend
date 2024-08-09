package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	db.InitDB(cfg)

	// Initialize Gin
	r := gin.Default()

	// Start the server
	r.Run(":" + cfg.AppPort)
}
