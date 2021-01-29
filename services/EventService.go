package services

import (
	"climbpass-matching-service/exceptions"
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"encoding/json"
	"fmt"
)

// IEventService interface for EventService
type IEventService interface {
	GetEventByID(id uint64) ([]byte, error)
	GetAllEvents() ([]byte, error)
	CreateEvent(Event models.EventModel) ([]byte, error)
	DeleteEventByID(id uint64) ([]byte, error)
	UpdateEventByID(id uint64, Event models.EventModel) ([]byte, error)
}

// EventService implementaion of interface
type EventService struct {
	repository repositories.IEventRepository
}

// NewEventService init
func NewEventService() EventService {
	repo := repositories.NewEventRepo()
	return EventService{repo}
}

// GetEventByID gets Events by id
func (service EventService) GetEventByID(id uint64) ([]byte, error) {
	Event := service.repository.GetEventByID(id)
	return json.Marshal(Event)
}

// GetAllEvents gets all Events
func (service EventService) GetAllEvents() ([]byte, error) {
	Events := service.repository.GetAllEvents()
	return json.Marshal(Events)
}

// CreateEvent creates a Event
func (service EventService) CreateEvent(Event models.EventModel) ([]byte, error) {
	existing := service.repository.CheckEventExists(Event)
	if existing.ID == 0 {
		newEvent := service.repository.CreateEvent(Event)
		return json.Marshal(newEvent)
	}
	return []byte(""), exceptions.EventExistsException()
}

// DeleteEventByID deletes if present, does nothing if not
func (service EventService) DeleteEventByID(id uint64) ([]byte, error) {
	service.repository.DeleteEventByID(id)
	return []byte(fmt.Sprintf("%v", id)), nil
}

// UpdateEventByID updates with new object
func (service EventService) UpdateEventByID(id uint64, Event models.EventModel) ([]byte, error) {
	Event.ID = id
	service.repository.UpdateEventByID(Event)
	return []byte(""), nil
}
