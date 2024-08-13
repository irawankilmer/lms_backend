package migrator

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/irawankilmer/lms_backend/internal/db/migrations"
	"gorm.io/gorm"
)

func MigrationRegistry() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		createMigration("20240814_create_users_table", migrations.Up_20240814_CreateUsersTable, migrations.Down_20240814_CreateUsersTable),
		createMigration("20240814_create_reset_tokens_table", migrations.Up_20240814_CreateResetTokensTable, migrations.Down_20240814_CreateResetTokensTable),
	}
}

func createMigration(id string, up func(tx *gorm.DB) error, down func(tx *gorm.DB) error) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: id,
		Migrate: func(tx *gorm.DB) error {
			return up(tx)
		},
		Rollback: func(tx *gorm.DB) error {
			return down(tx)
		},
	}
}
