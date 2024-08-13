package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Client struct {
	ID           uint   `gorm:"primaryKey"`
	ClientName   string `gorm:"not null;unique"`
	ClientID     string `gorm:"not null;unique"`
	ClientSecret string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// HashClientSecret hashes the client secret before storing it in the database
func (c *Client) HashClientSecret(secret string) error {
	hashedSecret, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.ClientSecret = string(hashedSecret)
	return nil
}

// CheckClientSecret checks if the provided secret matches the hashed secret
func (c *Client) CheckClientSecret(secret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.ClientSecret), []byte(secret))
	return err == nil
}
