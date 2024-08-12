package db

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	migrations := []*gormigrate.Migration{
		{
			ID: "20240812_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "20240812_create_reset_tokens_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.ResetToken{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("reset_tokens")
			},
		},
		// Tambahkan migrasi lainnya di sini
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	log.Println("Migration completed successfully")
	return nil
}
