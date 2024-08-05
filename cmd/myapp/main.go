package main

import (
	"log"

	_ "github.com/irawankilmer/lms_backend/docs"
	"github.com/irawankilmer/lms_backend/internal/config"
	"github.com/irawankilmer/lms_backend/internal/routes"
)

// @title LMS Backend API
// @version 1.0
// @description API Documentation for LMS Backend
// @host localhost:8080
// @BasePath /api

func main() {
	config.InitDB()
	router := routes.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
