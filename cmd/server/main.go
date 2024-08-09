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
	"github.com/irawankilmer/lms_backend/internal/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance := db.InitDB(cfg)

	// Run migrations
	if err := db.Migrate(dbInstance); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

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
			utils.JSON(c, 400, "Invalid input", nil, err.Error())
			return
		}

		token, err := authService.Login(input.Identifier, input.Password)
		if err != nil {
			utils.JSON(c, 401, "Authentication failed", nil, err.Error())
			return
		}

		utils.JSON(c, 200, "Login Successful", gin.H{"token": token}, "")
	})

	// Protected endpoint example
	r.GET("/protected", middleware.AuthMiddleware(cfg), func(c *gin.Context) {
		userID := c.MustGet("userID").(float64) // Assuming userID is of type float64
		utils.JSON(c, 200, "Authorization Successful", gin.H{"message": "Welcome", "user_id": userID}, "")
	})

	// CRUD Endpoints for User
	protected := r.Group("/users", middleware.AuthMiddleware(cfg))

	// Create User
	protected.POST("/", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.JSON(c, 400, "Invalid input", nil, err.Error())
			return
		}

		// Validate the user struct
		if err := validate.Struct(&user); err != nil {
			utils.JSON(c, 400, "Invalid input", nil, err.Error())
			return
		}

		if err := userService.CreateUser(&user); err != nil {
			utils.JSON(c, 500, "Failed to create user", nil, err.Error())
			return
		}

		utils.JSON(c, 201, "User created successfully", user, "")
	})

	// Start the server
	r.Run(":" + cfg.AppPort)
}
