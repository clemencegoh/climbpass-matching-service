package services

import (
	"climbpass-matching-service/exceptions"
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"encoding/json"
	"fmt"
)

// IUserService interface for UserService
type IUserService interface {
	GetUserByName(string) ([]byte, error)
	CreateUser(User models.User) (models.User, error)
	DeleteUserByID(id uint64) ([]byte, error)
	UpdateUserByID(id uint64, User models.User) ([]byte, error)
}

// UserService implementaion of interface
type UserService struct {
	repository repositories.IUserRepository
}

// NewUserService init
func NewUserService() UserService {
	repo := repositories.NewUserRepo()
	return UserService{repo}
}

// GetUserByName gets User by username to Userenticate
func (service UserService) GetUserByName(name string) ([]byte, error) {
	gym := service.repository.GetUserByName(name)
	return json.Marshal(gym)
}

// CreateUser creates a new user
func (service UserService) CreateUser(user models.User) (models.User, error) {
	existing := service.repository.GetUserByName(user.Name)
	if existing.ID == 0 {
		newUser := service.repository.CreateUser(user)
		return newUser, nil
	}
	return user, exceptions.UserAlreadyExists()
}

// DeleteUserByID deletes if present, does nothing if not
func (service UserService) DeleteUserByID(id uint64) ([]byte, error) {
	service.repository.DeleteUserByID(id)
	return []byte(fmt.Sprintf("%v", id)), nil
}

// UpdateUserByID updates with new object
func (service UserService) UpdateUserByID(id uint64, user models.User) ([]byte, error) {
	user.ID = id
	service.repository.UpdateUserByID(user)
	return []byte(""), nil
}
