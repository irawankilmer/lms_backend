package models

import "time"

type Permission struct {
	ID             uint   `gorm:"primaryKey"`
	PermissionName string `gorm:"unique; not null"`
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
