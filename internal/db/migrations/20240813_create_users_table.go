package migrations

import (
	"gorm.io/gorm"
	"time"
)

func Up_20240814_CreateUsersTable(tx *gorm.DB) error {
	type User struct {
		ID            uint   `gorm:"primaryKey"`
		Username      string `gorm:"unique;not null"`
		Email         string `gorm:"unique;not null"`
		Password      string `gorm:"not null"`
		EmailVerified time.Time
		LastLogin     time.Time
		LastActivity  time.Time
		CreatedAt     time.Time `gorm:"autoCreateTime"`
		UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	}
	return tx.AutoMigrate(&User{})
}

func Down_20240814_CreateUsersTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable("users")
}
