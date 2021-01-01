package models

import (
	"gorm.io/gorm"
)

// AuthUser for auth login
type AuthUser struct {
	gorm.Model
	ID       int `gorm:"primary_key;auto_increment;not_null"`
	Username string
	Password string

	User *User `gorm:"foreignKey:ID;"`
}
