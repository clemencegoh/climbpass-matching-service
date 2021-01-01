package repositories

import (
	"climbpass-matching-service/models"
	"gorm.io/gorm"
)

// IGymRepository exposed interface
type IGymRepository interface {
	GetGymByName(name string) models.GymModel
	GetAllGyms() []*models.GymModel
	CreateGym(gym models.GymModel) models.GymModel
	DeleteGymByID(id int)
	UpdateGymByID(gym models.GymModel) models.GymModel
}

// GymRepository struct
type GymRepository struct {
	db *gorm.DB
}

// NewGymRepo inits new GymRepo
func NewGymRepo() IGymRepository {
	db := Connect()
	return GymRepository{db}
}

// GetGymByName gets gym by specific name
func (repo GymRepository) GetGymByName(name string) models.GymModel {
	var gyms models.GymModel
	repo.db.First(&gyms, "name = ?", name)

	return gyms
}

// GetAllGyms gets all
func (repo GymRepository) GetAllGyms() []*models.GymModel {
	var gyms []*models.GymModel

	repo.db.Find(&gyms)

	return gyms
}

// CreateGym creates new gym if not already
func (repo GymRepository) CreateGym(gym models.GymModel) models.GymModel {
	repo.db.Create(&gym)
	return gym
}

// DeleteGymByID deletes a gym by its id
func (repo GymRepository) DeleteGymByID(id int) {
	repo.db.Delete(&models.GymModel{}, id)
}

// UpdateGymByID deletes a gym by its id
func (repo GymRepository) UpdateGymByID(gym models.GymModel) models.GymModel {
	repo.db.Save(&gym)
	return gym
}
