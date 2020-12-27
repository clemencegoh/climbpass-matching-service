package models

import (
	"gorm.io/gorm"
)

// GymModel for gyms
type GymModel struct {
	gorm.Model
	ID       int `gorm:"primary_key;auto_increment;not_null"`
	Name     string
	Location string
}
