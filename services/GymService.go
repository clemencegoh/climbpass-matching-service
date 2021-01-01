package services

import (
	"climbpass-matching-service/exceptions"
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"encoding/json"
)

// IGymService interface for GymService
type IGymService interface {
	GetGymByName(string) ([]byte, error)
	GetAllGyms() ([]byte, error)
	CreateGym(gym models.GymModel) ([]byte, error)
	DeleteGymByID(id int) ([]byte, error)
	UpdateGymByID(id int, gym models.GymModel) ([]byte, error)
}

// GymService implementaion of interface
type GymService struct {
	repository repositories.IGymRepository
}

// NewGymService init
func NewGymService() GymService {
	repo := repositories.NewGymRepo()
	return GymService{repo}
}

// GetGymByName gets gyms by name
func (service GymService) GetGymByName(name string) ([]byte, error) {
	gym := service.repository.GetGymByName(name)
	return json.Marshal(gym)
}

// GetAllGyms gets all gyms
func (service GymService) GetAllGyms() ([]byte, error) {
	gyms := service.repository.GetAllGyms()
	return json.Marshal(gyms)
}

// CreateGym creates a gym
func (service GymService) CreateGym(gym models.GymModel) ([]byte, error) {
	existing := service.repository.GetGymByName(gym.Name)
	if existing.ID == 0 {
		newGym := service.repository.CreateGym(gym)
		return json.Marshal(newGym)
	}
	return []byte(""), exceptions.GymExistsException(gym.Name)
}

// DeleteGymByID deletes if there, does nothing if not
func (service GymService) DeleteGymByID(id int) ([]byte, error) {
	service.repository.DeleteGymByID(id)
	return []byte(""), nil
}

// UpdateGymByID updates with new object
func (service GymService) UpdateGymByID(id int, gym models.GymModel) ([]byte, error) {
	gym.ID = id
	service.repository.UpdateGymByID(gym)
	return []byte(""), nil
}
