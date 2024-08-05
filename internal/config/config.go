package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		panic("Unsupported DB type")
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected")
}
