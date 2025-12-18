package service

import (
	"context"
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUserById(ctx context.Context, userID string) error
	UpdateUserById(ctx context.Context, userID string, user *model.User) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

// init a new UserService
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repo: repository}
}

// create user
func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.CreateUser(ctx, user)
}

// get user by id
func (s *userService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

// delete user by id
func (s *userService) DeleteUserById(ctx context.Context, userID string) error {
	return s.repo.DeleteUserById(ctx, userID)
}

// update user by id
func (s *userService) UpdateUserById(ctx context.Context, userID string, user *model.User) (*model.User, error) {
	return s.repo.UpdateUserById(ctx, user, userID)
}
