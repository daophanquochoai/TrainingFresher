package service

import (
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
	"testing"
)

// setup
func setup() service.UserService {
	repo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(repo)
	_, _ = userService.CreateUser(&model.User{
		ID:   "1",
		NAME: "Alice",
	})
	return userService
}

// Test_CreateUser_Succes tests the successful creation of a user.
func Test_CreateUser_Success(t *testing.T) {
	userService := setup()

	user := model.User{
		ID:   "2",
		NAME: "John Doe",
	}
	createdUser, err := userService.CreateUser(&user)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if createdUser.ID != user.ID || createdUser.NAME != user.NAME {
		t.Fatalf("expected user %v, got %v", user, createdUser)
	}
}

// Test_CreateUser_InvalidInput tests the creation of a user with invalid input.
func Test_CreateUser_InvalidInput(t *testing.T) {
	userService := setup()

	user := model.User{
		ID:   "",
		NAME: "John Doe",
	}
	_, err := userService.CreateUser(&user)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != service.ErrValidInput {
		t.Fatalf("expected ErrValidInput, got %v", err)
	}
}

// Test_GetUserByID_Success tests the successful retrieval of a user by ID.
func Test_GetUserByID_Success(t *testing.T) {
	userService := setup()

	userID := "1"
	user, err := userService.GetUserByID(userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != userID {
		t.Fatalf("expected user ID %v, got %v", userID, user.ID)
	}
}

// Test_GetUserByID_InvalidInput tests the retrieval of a user with invalid ID.
func Test_GetUserByID_InvalidInput(t *testing.T) {
	userService := setup()

	userID := ""
	_, err := userService.GetUserByID(userID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != service.ErrValidInput {
		t.Fatalf("expected ErrValidInput, got %v", err)
	}
}
