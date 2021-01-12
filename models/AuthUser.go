package models

import (
	"gorm.io/gorm"
)

// AuthUser for auth login
type AuthUser struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Username string
	Password string

	User User `gorm:"foreignKey:ID"`
}
