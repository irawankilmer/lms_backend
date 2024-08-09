package db

import (
	"log"

	"github.com/irawankilmer/lms_backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize database connection
func InitDB(cfg *config.Config) *gorm.DB {
	var err error

	switch cfg.DBType {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	case "mysql":
		DB, err = gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
	default:
		log.Fatalf("Unsupported DB type: %s", cfg.DBType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return DB
}
