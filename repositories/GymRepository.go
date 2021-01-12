package repositories

import (
	"climbpass-matching-service/models"

	"gorm.io/gorm"
)

// IGymRepository exposed interface
type IGymRepository interface {
	GetGymByName(name string) models.GymProfile
	GetAllGyms() []*models.GymProfile
	CreateGym(gym models.GymProfile) models.GymProfile
	DeleteGymByID(id uint64)
	UpdateGymByID(gym models.GymProfile) models.GymProfile
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
func (repo GymRepository) GetGymByName(name string) models.GymProfile {
	var gyms models.GymProfile
	repo.db.First(&gyms, "name = ?", name)

	return gyms
}

// GetAllGyms gets all
func (repo GymRepository) GetAllGyms() []*models.GymProfile {
	var gyms []*models.GymProfile

	repo.db.Find(&gyms)

	return gyms
}

// CreateGym creates new gym if not already
func (repo GymRepository) CreateGym(gym models.GymProfile) models.GymProfile {
	repo.db.Create(&gym)
	return gym
}

// DeleteGymByID deletes a gym by its id
func (repo GymRepository) DeleteGymByID(id uint64) {
	repo.db.Delete(&models.GymProfile{}, id)
}

// UpdateGymByID updates a gym by its id
func (repo GymRepository) UpdateGymByID(gym models.GymProfile) models.GymProfile {
	repo.db.Save(&gym)
	return gym
}
