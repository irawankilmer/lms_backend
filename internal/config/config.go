package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/irawankilmer/lms_backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DSN")

	switch dbType {
	case "postgres":
		DB, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // Disable prepared statements
		}), &gorm.Config{
			PrepareStmt: false, // Ensure prepared statements are not used
		})
	case "mysql":
		DB, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			PrepareStmt: false, // Ensure prepared statements are not used
		})
	default:
		panic("Unsupported DB type")
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB from GORM DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// AutoMigrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.Client{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	fmt.Println("Database connected")
}
