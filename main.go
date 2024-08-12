package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	_ "github.com/irawankilmer/lms_backend/docs"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/handler"
	"github.com/irawankilmer/lms_backend/internal/router"
	"github.com/irawankilmer/lms_backend/internal/service"
)

var (
	app *gin.Engine
	cfg *config.Config
)

func init() {
	// Load configuration
	var err error
	cfg, err = config.LoadConfig()
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
	app = router.SetupRouter(cfg)
}

func main() {
	// Determine environment (development or production)
	env := os.Getenv("APP_ENV")

	if env == "production" {
		// Running on Vercel
		http.HandleFunc("/", Handler)
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080" // Default port jika tidak ada
		}
		log.Printf("Starting server on port %s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {
		// Running locally
		log.Printf("Starting server on port %s", cfg.AppPort)
		if err := app.Run(":" + cfg.AppPort); err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}
}

// Handler for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
