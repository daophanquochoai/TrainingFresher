package service

import (
	"errors"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
)

var ErrValidInput = errors.New("Invalid input")

type UserService interface {
	CreateUser(u *model.User) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService initializes a new UserService with the given UserRepository.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

// Create user after validating input.
func (s *userService) CreateUser(u *model.User) (*model.User, error) {
	if u.ID == "" || u.NAME == "" {
		return nil, ErrValidInput
	}
	return s.repo.Create(u)
}

// Get user by ID.
func (s *userService) GetUserByID(id string) (*model.User, error) {
	if id == "" {
		return nil, ErrValidInput
	}
	return s.repo.GetByID(id)
}
