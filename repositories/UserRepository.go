package repositories

import (
	"climbpass-matching-service/models"

	"gorm.io/gorm"
)

// IUserRepository exposed interface
type IUserRepository interface {
	GetUserByName(name string) models.User
	CreateUser(user models.User) models.User
	DeleteUserByID(id uint64)
	UpdateUserByID(user models.User) models.User
}

// UserRepository struct
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepo inits new AuthRepo
func NewUserRepo() IUserRepository {
	db := Connect()
	return UserRepository{db}
}

// GetUserByName gets by name
func (repo UserRepository) GetUserByName(name string) models.User {
	var user models.User
	repo.db.First(&user, "name = ?", name)

	return user
}

// CreateUser creates new user
func (repo UserRepository) CreateUser(user models.User) models.User {
	repo.db.Create(&user)
	return user
}

// DeleteUserByID deletes by id
func (repo UserRepository) DeleteUserByID(id uint64) {
	repo.db.Delete(&models.User{}, id)
}

// UpdateUserByID updates a user
func (repo UserRepository) UpdateUserByID(user models.User) models.User {
	repo.db.Save(&user)
	return user
}
