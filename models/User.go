package models

import (
	"gorm.io/gorm"
)

// User for normal climbing users
type User struct {
	gorm.Model
	ID int `gorm:"primary_key;auto_increment;not_null"`

	Name          string
	ClimbingGrade string

	Gyms []GymProfile `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
