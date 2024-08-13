package migrator

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, MigrationRegistry())

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	log.Println("Migration completed successfully")
	return nil
}
