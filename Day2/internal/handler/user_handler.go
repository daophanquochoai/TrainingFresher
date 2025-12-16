package handler

import (
	"encoding/json"
	"go-db-demo/internal/model"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

// init a new UserHandler
func NewHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// create user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//clean up it
	defer r.Body.Close()

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := model.Response{
			Message: "Invalid request payload",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusBadRequest, response)
		return
	}
	userSaved, error := h.service.CreateUser(&user)
	if error != nil {
		response := model.Response{
			Message: error.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Data:    userSaved,
		Message: "User created successfully",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusCreated, response)

}

// get user by id
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userID := r.URL.Path[len("/users/"):]

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		response := model.Response{
			Message: "Internal server error",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	if user == nil {
		response := model.Response{
			Message: "User not found",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusNotFound, response)
		return
	}
	response := model.Response{
		Data:    user,
		Message: "User fetched successfully",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusOK, response)
}
