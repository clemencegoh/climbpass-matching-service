package models

import (
	"gorm.io/gorm"
)

// GymProfile for gyms
type GymProfile struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key;auto_increment;" json:"id;"`
	Name     string
	Location string
}
