package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/middleware"
	"github.com/irawankilmer/lms_backend/internal/models"
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

	// Initialize Service
	authService := &service.AuthService{DB: dbInstance, Config: cfg}
	userService := &service.UserService{DB: dbInstance}

	// Initialize validator
	validate := validator.New()
	// Register custom validation for no spaces in username
	validate.RegisterValidation("nospaces", func(fl validator.FieldLevel) bool {
		return !strings.Contains(fl.Field().String(), " ")
	})

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

	// Protected endpoint example
	r.GET("/protected", middleware.AuthMiddleware(cfg), func(c *gin.Context) {
		userID := c.MustGet("userID").(float64) // Assuming userID is of type float64
		c.JSON(200, gin.H{"message": "Welcome!", "user_id": userID})
	})

	// CRUD Endpoints for User
	protected := r.Group("/users", middleware.AuthMiddleware(cfg))

	// Create User
	protected.POST("/", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Validate the user struct
		if err := validate.Struct(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := userService.CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, user)
	})

	// Start the server
	r.Run(":" + cfg.AppPort)
}
