package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance := db.InitDB(cfg)

	// Initialize Gin
	r := gin.Default()

	// Initialize AuthService
	authService := &service.AuthService{DB: dbInstance, Config: cfg}

	// Login endpoint
	r.POST("/login", func(c *gin.Context) {
		var input struct {
			Identifier string `json:"identifier" binding:"required"`
			Password   string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := authService.Login(input.Identifier, input.Password)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": token})
	})

	// Start the server
	r.Run(":" + cfg.AppPort)
}
