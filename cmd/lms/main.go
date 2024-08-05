package main

import (
	"github.com/irawankilmer/lms_backend/internal/config"
	"github.com/irawankilmer/lms_backend/internal/models"
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
}
