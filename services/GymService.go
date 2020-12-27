package services

import (
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"encoding/json"
)

type IGymService interface {
	GetGymByName(string) ([]byte, error)
	GetAllGyms() ([]byte, error)
	CreateGym(gym models.GymModel) ([]byte, error)
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

func (service GymService) GetGymByName(name string) ([]byte, error) {
	gym := service.repository.GetGymByName(name)
	return json.Marshal(gym)
}

func (service GymService) GetAllGyms() ([]byte, error) {
	gyms := service.repository.GetAllGyms()
	return json.Marshal(gyms)
}

func (service GymService) CreateGym(gym models.GymModel) ([]byte, error) {
	service.repository.CreateGym(gym)
	return json.Marshal(gym)
}
