package models

import "time"

type Client struct {
	ClientID     uint   `gorm:"primaryKey"`
	ClientName   string `gorm:"unique;not null"`
	ClientSecret string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
