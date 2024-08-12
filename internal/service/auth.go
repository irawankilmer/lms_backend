package service

import (
	"errors"
	"fmt"
	"time"

	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/irawankilmer/lms_backend/config"
	"github.com/irawankilmer/lms_backend/internal/models"
	"github.com/irawankilmer/lms_backend/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB           *gorm.DB
	Config       *config.Config
	EmailService *EmailService
}

// NewAuthService creates a new instance of AuthService with the given dependencies.
func NewAuthService(db *gorm.DB, config *config.Config, emailService *EmailService) *AuthService {
	return &AuthService{
		DB:           db,
		Config:       config,
		EmailService: emailService,
	}
}

// Login authenticates a user and returns a JWT token.
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

// GeneratePasswordResetToken generates a reset token and stores it in the database.
func (s *AuthService) GeneratePasswordResetToken(email string) (string, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return "", err
	}
	log.Printf("Generated token: %s", token)

	expiresAt := time.Now().Add(1 * time.Hour) // Token valid for 1 hour

	resetToken := models.ResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	if err := s.DB.Create(&resetToken).Error; err != nil {
		return "", err
	}

	return token, nil
}

// ResetPassword resets the user's password using the provided token.
func (s *AuthService) ResetPassword(token, newPassword string) error {
	var resetToken models.ResetToken
	if err := s.DB.Where("token = ?", token).First(&resetToken).Error; err != nil {
		return errors.New("invalid or expired token")
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return errors.New("token has expired")
	}

	var user models.User
	if err := s.DB.First(&user, resetToken.UserID).Error; err != nil {
		return err
	}

	if err := user.HashPassword(newPassword); err != nil {
		return err
	}

	if err := s.DB.Save(&user).Error; err != nil {
		return err
	}

	// Optionally, delete the used token
	s.DB.Delete(&resetToken)

	return nil
}

// SendPasswordResetEmail sends a password reset email to the user.
func (s *AuthService) SendPasswordResetEmail(email, token string) error {
	resetURL := fmt.Sprintf("https://yourdomain.com/password-reset?token=%s", token)
	subject := "Password Reset Request"
	body := fmt.Sprintf("Click the link to reset your password: %s", resetURL)

	err := s.EmailService.SendEmail(email, subject, body)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
