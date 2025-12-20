package service

import (
	"context"
	"errors"
	"product/internal/model"
	"product/internal/repository"
)

var ErrValidate = " is not valid"

type ProductService interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
}

type productService struct {
	r repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		r: r,
	}
}

func (s *productService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	if id == "" {
		return nil, errors.New("Id" + ErrValidate)
	}
	return s.r.GetProduct(ctx, id)
}

func (s *productService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	if product.Name == "" {
		return nil, errors.New("Name" + ErrValidate)
	}
	if product.Price <= 0 {
		return nil, errors.New("Price" + ErrValidate)
	}
	return s.r.CreateProduct(ctx, product)
}
