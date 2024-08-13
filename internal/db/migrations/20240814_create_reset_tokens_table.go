package migrations

import (
	"gorm.io/gorm"
	"time"
)

func Up_20240814_CreateResetTokensTable(tx *gorm.DB) error {
	type ResetToken struct {
		ID        uint      `gorm:"primaryKey"`
		UserID    uint      `gorm:"not null"`
		Token     string    `gorm:"unique;not null"`
		ExpiresAt time.Time `gorm:"not null"`
		CreatedAt time.Time
	}
	return tx.AutoMigrate(&ResetToken{})
}

func Down_20240814_CreateResetTokensTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable("reset_tokens")
}
