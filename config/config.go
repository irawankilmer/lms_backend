package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config structure to hold configuration values
type Config struct {
	AppPort     string
	AppEnv      string
	DBType      string
	PostgresDSN string
	MySQLDSN    string
	JwtSecret   string
	RedisAddr   string
	RedisDB     int
	RedisPass   string
}

// Initialize and return configuration
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
		return nil, err
	}

	config := &Config{
		AppPort:     viper.GetString("APP_PORT"),
		AppEnv:      viper.GetString("APP_ENV"),
		DBType:      viper.GetString("DB_TYPE"),
		PostgresDSN: viper.GetString("POSTGRES_DSN"),
		MySQLDSN:    viper.GetString("MYSQL_DSN"),
		JwtSecret:   viper.GetString("JWT_SECRET"),
		RedisAddr:   viper.GetString("REDIS_ADDR"),
		RedisDB:     viper.GetInt("REDIS_DB"),
		RedisPass:   viper.GetString("REDIS_PASSWORD"),
	}

	return config, nil
}
