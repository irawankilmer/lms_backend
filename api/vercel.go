package api

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	_ "github.com/irawankilmer/lms_backend/docs"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/handler"
	"github.com/irawankilmer/lms_backend/internal/router"
	"github.com/irawankilmer/lms_backend/internal/service"
	"log"
	"net/http"
)

var (
	app *gin.Engine
)

// @title MyApp API
// @version 1.0
// @description This is a sample server for MyApp.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Initialize the database
	dbInstance := db.InitDB(cfg)

	// Initialize services with dbInstance
	userService := &service.UserService{DB: dbInstance}
	authService := &service.AuthService{DB: dbInstance, Config: cfg}
	handler.SetServices(userService, authService)

	// Set up routes
	r := router.SetupRouter(cfg)

	// Start server
	r.Run(":" + cfg.AppPort)
}

// vercel net/http entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}