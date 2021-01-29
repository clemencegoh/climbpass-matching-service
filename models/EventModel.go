package models

import (
	"gorm.io/gorm"
)

// EventModel model
type EventModel struct {
	gorm.Model
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ChatroomID    string
	Name          string
	Description   string
	EventDate     string
	Organizer     User
	HaveMultipass bool
	Members       []User
}
