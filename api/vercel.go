package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/db"
	"github.com/irawankilmer/lms_backend/internal/handler"
	"github.com/irawankilmer/lms_backend/internal/router"
	"github.com/irawankilmer/lms_backend/internal/service"
)

var app *gin.Engine

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
	app = router.SetupRouter(cfg)
}

// vercel net/http entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request:", r.URL.Path)
	app.ServeHTTP(w, r)
}
