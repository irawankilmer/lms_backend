package main

import (
	handler "github.com/irawankilmer/lms_backend/api"
	"github.com/irawankilmer/lms_backend/internal/config"
	"github.com/irawankilmer/lms_backend/internal/models"
	"net/http"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
	)

	http.HandleFunc("/api", handler.Handler)
	http.ListenAndServe(":8080", nil)
}
