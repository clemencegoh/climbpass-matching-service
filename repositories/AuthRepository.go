package repositories

import (
	"climbpass-matching-service/models"

	"gorm.io/gorm"
)

// IAuthRepository exposed interface
type IAuthRepository interface {
	GetAuthByUsername(name string) models.AuthUser
	CreateAuth(Auth models.AuthUser) models.AuthUser
	DeleteAuthByID(id uint64)
	UpdateAuthByID(Auth models.AuthUser) models.AuthUser
}

// AuthRepository struct
type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepo inits new AuthRepo
func NewAuthRepo() IAuthRepository {
	db := Connect()
	return AuthRepository{db}
}

// GetAuthByUsername gets Auth by username
func (repo AuthRepository) GetAuthByUsername(username string) models.AuthUser {
	var auths models.AuthUser
	repo.db.First(&auths, "username = ?", username)

	return auths
}

// CreateAuth creates new Auth for login
func (repo AuthRepository) CreateAuth(Auth models.AuthUser) models.AuthUser {
	repo.db.Create(&Auth)
	return Auth
}

// DeleteAuthByID deletes a auth user
func (repo AuthRepository) DeleteAuthByID(id uint64) {
	repo.db.Delete(&models.AuthUser{}, id)
}

// UpdateAuthByID updates a Auth user by its id
func (repo AuthRepository) UpdateAuthByID(Auth models.AuthUser) models.AuthUser {
	repo.db.Save(&Auth)
	return Auth
}
