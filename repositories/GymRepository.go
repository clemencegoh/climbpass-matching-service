package repositories

import (
	"climbpass-matching-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GymRepository struct{}

func (repo *GymRepository) GetGymByName(name string) models.GymModel {
	var gyms models.GymModel

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db.First(&gyms, "name = ?", name)
	return gyms
}
