package seeder

import "gorm.io/gorm"

// Seed runs all the seed functions
func Seed(db *gorm.DB) error {
	if err := SeedUsers(db); err != nil {
		return err
	}

	if err := SeedClients(db); err != nil {
		return err
	}
	// Tambahkan seeder lainnya disini

	return nil
}
