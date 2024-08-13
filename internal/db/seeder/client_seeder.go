package seeder

import (
	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
	"log"
)

func SeedClients(db *gorm.DB) error {
	clients := []models.Client{
		{
			ClientName: "WebApp",
			ClientID:   "web-app-id",
		},
		{
			ClientName: "MobileApp",
			ClientID:   "mobile-app-id",
		},
	}

	// Hash the client secrets and save the clients
	for _, client := range clients {
		secret := "supersecretkey" // Replace with the actual secret you want to use
		if err := client.HashClientSecret(secret); err != nil {
			return err
		}
		if err := db.Where("client_id = ?", client.ClientID).FirstOrCreate(&client).Error; err != nil {
			return err
		}
	}

	log.Println("Clients seeded successfully")
	return nil
}
