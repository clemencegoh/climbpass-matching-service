package models

import (
	"gorm.io/gorm"
)

// GymModel for gyms
type GymModel struct {
	gorm.Model
	ID       int
	Name     string
	Location string
}
