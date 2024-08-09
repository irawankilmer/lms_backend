package db

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Define migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230810_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		// Add more migrations here
	})

	// Run migrations
	return m.Migrate()
}
