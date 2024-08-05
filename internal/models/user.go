package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique; not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique"`
	FullName  string
	Roles     []Role `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
