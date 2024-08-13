package migrations

import (
	"gorm.io/gorm"
	"time"
)

func Up_20240814_CreateClientsTable(tx *gorm.DB) error {
	type Client struct {
		ID           uint   `gorm:"primaryKey"`
		ClientName   string `gorm:"not null;unique"`
		ClientID     string `gorm:"not null;unique"`
		ClientSecret string `gorm:"not null"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	return tx.AutoMigrate(&Client{})
}

func Down_20240814_CreateClientsTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable("clients")
}
