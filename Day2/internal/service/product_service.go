package service

import (
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	GetProductById(productId string) (*model.Product, error)
}

type productService struct {
	repoInstance repository.ProductRepository
}

// init
func NewProductService(repoParam repository.ProductRepository) ProductService {
	return &productService{repoInstance: repoParam}
}

// create product
func (r *productService) CreateProduct(product *model.Product) (*model.Product, error) {
	return r.repoInstance.CreateProduct(product)
}

// get product by id
func (r *productService) GetProductById(productId string) (*model.Product, error) {
	return r.repoInstance.GetProductById(productId)
}
