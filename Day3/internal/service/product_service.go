package service

import (
	"context"
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetProductById(ctx context.Context, productId string) (*model.Product, error)
	UpdateProductById(ctx context.Context, productId string, product *model.Product) (*model.Product, error)
	DeleteProductById(ctx context.Context, productId string) error
}

type productService struct {
	repoInstance repository.ProductRepository
}

// init
func NewProductService(repoParam repository.ProductRepository) ProductService {
	return &productService{repoInstance: repoParam}
}

// create product
func (r *productService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	return r.repoInstance.CreateProduct(ctx, product)
}

// get product by id
func (r *productService) GetProductById(ctx context.Context, productId string) (*model.Product, error) {
	return r.repoInstance.GetProductById(ctx, productId)
}

// update product
func (r *productService) UpdateProductById(ctx context.Context, productId string, product *model.Product) (*model.Product, error) {
	return r.repoInstance.UpdateProductById(ctx, productId, product)
}

// delete product
func (r *productService) DeleteProductById(ctx context.Context, productId string) error {
	return r.repoInstance.DeleteProductById(ctx, productId)
}
