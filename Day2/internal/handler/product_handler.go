package handler

import (
	"encoding/json"
	"go-db-demo/internal/model"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"net/http"
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
	productSaved, error := s.serviceInstance.CreateProduct(&product)
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
	productId := r.URL.Path[len("/products/"):]
	product, err := s.serviceInstance.GetProductById(productId)
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
