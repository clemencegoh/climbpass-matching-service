package models

import (
	"gorm.io/gorm"
)

// GymProfile for gyms
type GymProfile struct {
	gorm.Model
	ID       int `gorm:"primary_key;auto_increment;not_null"`
	Name     string
	Location string
}
