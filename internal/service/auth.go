package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
)

// AuthService represents the service for authentication
type AuthService struct {
	DB     *gorm.DB
	Config *config.Config
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(identifier, password string) (string, error) {
	var user models.User

	// Search by email or username
	if err := s.DB.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Check password
	if !user.CheckPassword(password) {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Config.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
