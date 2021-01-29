package repositories

import (
	"climbpass-matching-service/models"

	"gorm.io/gorm"
)

// IEventRepository exposed interface
type IEventRepository interface {
	GetEventByID(id uint64) models.EventModel
	CheckEventExists(Event models.EventModel) models.EventModel
	GetAllEvents() []*models.EventModel
	CreateEvent(Event models.EventModel) models.EventModel
	DeleteEventByID(id uint64)
	UpdateEventByID(Event models.EventModel) models.EventModel
}

// EventRepository struct
type EventRepository struct {
	db *gorm.DB
}

// NewEventRepo inits new EventRepo
func NewEventRepo() IEventRepository {
	db := Connect()
	return EventRepository{db}
}

// GetEventByID gets Event by ID
func (repo EventRepository) GetEventByID(id uint64) models.EventModel {
	var Events models.EventModel
	repo.db.First(&Events, "id = ?", id)

	return Events
}

// CheckEventExists gets Event by chat ID (for websocket)
func (repo EventRepository) CheckEventExists(Model models.EventModel) models.EventModel {
	var Events models.EventModel
	repo.db.Where("name = ? AND organizer.name = ?", Model.Name, Model.Organizer.Name).First(&Events)

	return Events
}

// GetAllEvents gets all
func (repo EventRepository) GetAllEvents() []*models.EventModel {
	var Events []*models.EventModel

	repo.db.Find(&Events)

	return Events
}

// CreateEvent creates new Event if not already
func (repo EventRepository) CreateEvent(Event models.EventModel) models.EventModel {
	repo.db.Create(&Event)
	return Event
}

// DeleteEventByID deletes a Event by its id
func (repo EventRepository) DeleteEventByID(id uint64) {
	repo.db.Delete(&models.EventModel{}, id)
}

// UpdateEventByID updates a Event by its id
func (repo EventRepository) UpdateEventByID(Event models.EventModel) models.EventModel {
	repo.db.Save(&Event)
	return Event
}
