package handler

import (
	"encoding/json"
	"errors"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
	"go-rest-api/internal/utils"
	"net/http"
	"strings"
)

// declare struct
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler initializes a new UserHandler with the given UserService.
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

// Handle creating a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// clean up body after reading
	defer r.Body.Close()

	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, model.Response{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}
	createUser, err := h.userService.CreateUser(&u)
	if err != nil {
		if errors.Is(err, service.ErrValidInput) {
			utils.WriteJson(w, http.StatusInternalServerError, model.Response{
				Status:  "error",
				Message: "Invalid user ID",
			})
			return
		}
		utils.WriteJson(w, http.StatusInternalServerError, model.Response{
			Status:  "error",
			Message: "Server error",
		})
		return
	}

	utils.WriteJson(w, http.StatusCreated, model.Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    createUser,
	})
}

// handle getting a user by ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, service.ErrValidInput) {
			utils.WriteJson(w, http.StatusInternalServerError, model.Response{
				Status:  "error",
				Message: "Invalid user ID",
			})
			return
		}
		if errors.Is(err, repository.ErrNotFound) {
			utils.WriteJson(w, http.StatusNotFound, model.Response{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		utils.WriteJson(w, http.StatusInternalServerError, model.Response{
			Status:  "error",
			Message: "Server error",
		})
	} else {
		utils.WriteJson(w, http.StatusOK, model.Response{
			Status:  "success",
			Message: "User retrieved successfully",
			Data:    user,
		})
	}
}
