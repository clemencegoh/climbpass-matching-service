package models

import (
	"gorm.io/gorm"
)

// User for normal climbing users
type User struct {
	gorm.Model
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name          string
	ClimbingGrade string
}
