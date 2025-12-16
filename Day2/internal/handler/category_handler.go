package handler

import (
	"encoding/json"
	"go-db-demo/internal/model"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"net/http"
)

type CategoryHandler struct {
	serviceInstance service.CategoryService
}

// init
func NewCategoryHandler(serviceParam service.CategoryService) *CategoryHandler {
	return &CategoryHandler{serviceInstance: serviceParam}
}

// create category
func (c *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response := model.Response{
			Message: "Method invalid allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	defer r.Body.Close()

	var category model.Categories
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	categorySaved, error := c.serviceInstance.CreateCategory(&category)
	if error != nil {
		response := model.Response{
			Message: error.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Data:    categorySaved,
		Message: "User created successfully",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusCreated, response)
}

// get category by id
func (c *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		response := model.Response{
			Message: "Method invalid allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}
	categoryId := r.URL.Path[len("/categories/"):]
	// get category
	category, err := c.serviceInstance.GetCategoryById(categoryId)
	if err != nil {
		response := model.Response{
			Message: "Internal server error",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	if category == nil {
		response := model.Response{
			Message: "User not found",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusNotFound, response)
		return
	}
	response := model.Response{
		Data:    category,
		Message: "User fetched successfully",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusOK, response)
}
