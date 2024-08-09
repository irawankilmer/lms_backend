package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents the user model
type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	EmailVerified *time.Time
	LastLogin     *time.Time
	LastActivity  *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// HashPassword hashes a plain text password
func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the password is correct
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
