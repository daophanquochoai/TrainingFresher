package handler

import (
	"context"
	"encoding/json"
	"go-db-demo/internal/model"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"net/http"
	"time"
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

	// timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

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
	categorySaved, error := c.serviceInstance.CreateCategory(ctx, &category)
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	categoryId := r.URL.Path[len("/categories/"):]
	// get category
	category, err := c.serviceInstance.GetCategoryById(ctx, categoryId)
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

// update category by id
func (h *CategoryHandler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	// timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	categoryId := r.URL.Path[len("/categories/update/"):]

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
	categoryReturn, error := h.serviceInstance.UpdateCategoryById(ctx, categoryId, &category)
	if error != nil {
		response := model.Response{
			Message: error.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Category updated successful",
		Status:  "success",
		Data:    categoryReturn,
	}
	utils.WriteJson(w, http.StatusOK, response)
}

// delete category
func (h *CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
	}

	// timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	categoryId := r.URL.Path[len("/categories/update/"):]

	err := h.serviceInstance.DeleteCategoryById(ctx, categoryId)
	if err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Delete successful",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusOK, response)
}

// add product into category
func (s *CategoryHandler) AddProductIntoCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	// timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	var categoryRequest model.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	categoryProduct, err := s.serviceInstance.AddProductIntoCategory(ctx, &categoryRequest)
	if err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Message: "Add product into category successfully",
		Status:  "success",
		Data:    categoryProduct,
	}
	utils.WriteJson(w, http.StatusOK, response)
}
