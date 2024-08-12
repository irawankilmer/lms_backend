package seeder

import (
	"log"
	"time"

	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
)

// SeedUsers seeds the users table with initial data
func SeedUsers(db *gorm.DB) error {
	now := time.Now()
	users := []models.User{
		{
			Username:      "admin",
			Email:         "admin@example.com",
			EmailVerified: &now,
			LastLogin:     &now,
			LastActivity:  &now,
		},
		{
			Username:      "admin2",
			Email:         "admin2@example.com",
			EmailVerified: &now,
			LastLogin:     &now,
			LastActivity:  &now,
		},
	}

	for _, user := range users {
		if err := db.Where("username = ?", user.Username).FirstOrCreate(&user).Error; err != nil {
			return err
		}
	}

	log.Println("Users seeded successfully")
	return nil
}
