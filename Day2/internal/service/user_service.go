package service

import (
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type UserService interface {
	GetUserByID(userID string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	DeleteUserById(userID string) error
	UpdateUserById(userID string, user *model.User) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

// init a new UserService
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repo: repository}
}

// create user
func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return s.repo.CreateUser(user)
}

// get user by id
func (s *userService) GetUserByID(userID string) (*model.User, error) {
	return s.repo.GetUserByID(userID)
}

// delete user by id
func (s *userService) DeleteUserById(userID string) error {
	return s.repo.DeleteUserById(userID)
}

// update user by id
func (s *userService) UpdateUserById(userID string, user *model.User) (*model.User, error) {
	return s.repo.UpdateUserById(user, userID)
}
