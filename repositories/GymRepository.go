package repositories

import (
	"climbpass-matching-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GymRepository exposed interface
type IGymRepository interface {
	GetGymByName(name string) models.GymModel
	GetAllGyms() []*models.GymModel
	CreateGym(gym models.GymModel)
}

// GymRepository struct
type GymRepository struct {
	db *gorm.DB
}

// NewGymRepo inits new GymRepo
func NewGymRepo() IGymRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
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

func (repo GymRepository) CreateGym(gym models.GymModel) {
	repo.db.Create(&gym)
}
