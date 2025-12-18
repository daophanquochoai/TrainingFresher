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

type ProductHandler struct {
	serviceInstance service.ProductService
}

// init
func NewProductHandler(serviceParam service.ProductService) *ProductHandler {
	return &ProductHandler{serviceInstance: serviceParam}
}

// create product
func (s *ProductHandler) CreateProductHanler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	// clean up
	defer r.Body.Close()

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	productSaved, error := s.serviceInstance.CreateProduct(ctx, &product)
	if error != nil {
		response := model.Response{
			Message: error.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Message: "Product created successfully",
		Data:    productSaved,
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusCreated, response)
}

// get product by id
func (s *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	productId := r.URL.Path[len("/products/"):]
	product, err := s.serviceInstance.GetProductById(ctx, productId)
	if err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Product found successfully",
		Data:    product,
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusOK, response)
}

// update product by id
func (s *ProductHandler) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		response := model.Response{
			Message: "Method not allowed",
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	productId := r.URL.Path[len("/products/update/"):]
	defer r.Body.Close()
	var productUpdate model.Product
	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		response := model.Response{
			Message: err.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	productUpdated, er := s.serviceInstance.UpdateProductById(ctx, productId, &productUpdate)
	if er != nil {
		response := model.Response{
			Message: er.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Product updated successful",
		Status:  "success",
		Data:    productUpdated,
	}
	utils.WriteJson(w, http.StatusOK, response)
	return
}

// delete product by id
func (s *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
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

	productId := r.URL.Path[len("/products/update/"):]
	er := s.serviceInstance.DeleteProductById(ctx, productId)
	if er != nil {
		response := model.Response{
			Message: er.Error(),
			Status:  "error",
		}
		utils.WriteJson(w, http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Delete product successfully",
		Status:  "success",
	}
	utils.WriteJson(w, http.StatusOK, response)
}
