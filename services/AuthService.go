package services

import (
	"climbpass-matching-service/configs/auth"
	"climbpass-matching-service/exceptions"
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"encoding/json"
	"fmt"
)

// IAuthService interface for AuthService
type IAuthService interface {
	GetAuthByUsername(string) ([]byte, error)
	CreateAuth(auth models.AuthUser) ([]byte, error)
	AuthenticateUser(auth models.AuthUser) ([]byte, error)
	DeleteAuthByID(id uint64) ([]byte, error)
	UpdateAuthByID(id uint64, auth models.AuthUser) ([]byte, error)
}

// AuthService implementaion of interface
type AuthService struct {
	repository repositories.IAuthRepository
}

// NewAuthService init
func NewAuthService() AuthService {
	repo := repositories.NewAuthRepo()
	return AuthService{repo}
}

// GetAuthByUsername gets auth by username to authenticate
func (service AuthService) GetAuthByUsername(username string) ([]byte, error) {
	gym := service.repository.GetAuthByUsername(username)
	return json.Marshal(gym)
}

// AuthenticateUser authenticates a user
func (service AuthService) AuthenticateUser(authUser models.AuthUser) ([]byte, error) {
	existing := service.repository.GetAuthByUsername(authUser.Username)
	if existing.ID == 0 || existing.Password != authUser.Password {
		// User not found
		return []byte(""), exceptions.UserNotFoundException()
	}
	token, err := auth.CreateToken(authUser.ID)
	if err != nil {
		return []byte(""), exceptions.UserNotFoundException()
	}
	return []byte(token), nil
}

// CreateAuth creates a new user
func (service AuthService) CreateAuth(user models.AuthUser) ([]byte, error) {
	existing := service.repository.GetAuthByUsername(user.Username)
	if existing.ID == 0 {
		newAuth := service.repository.CreateAuth(user)
		return json.Marshal(newAuth)
	}
	return []byte(""), exceptions.UserAlreadyExists()
}

// DeleteAuthByID deletes if present, does nothing if not
func (service AuthService) DeleteAuthByID(id uint64) ([]byte, error) {
	service.repository.DeleteAuthByID(id)
	return []byte(fmt.Sprintf("%v", id)), nil
}

// UpdateAuthByID updates with new object
func (service AuthService) UpdateAuthByID(id uint64, user models.AuthUser) ([]byte, error) {
	user.ID = id
	service.repository.UpdateAuthByID(user)
	return []byte(""), nil
}
