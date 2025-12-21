package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"product/internal/dto"
	"product/internal/model"
	"product/internal/repository"
)

var ErrValidate = " is not valid"

type ProductService interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.CreateProductRequest) (*model.Product, error)
	GetProductByFilter(ctx context.Context, filter *dto.PageRequest) (*dto.PageResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProductByCategoryId(ctx context.Context, categoryId uuid.UUID) ([]*model.Product, error)
}

type productService struct {
	r repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		r: r,
	}
}

func (s *productService) GetProductByCategoryId(ctx context.Context, categoryId uuid.UUID) ([]*model.Product, error) {
	if categoryId == uuid.Nil {
		return nil, errors.New("CategoryId " + ErrValidate)
	}
	return s.r.GetProductByCategoryId(ctx, categoryId)
}

func (s *productService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	if id == "" {
		return nil, errors.New("Id" + ErrValidate)
	}
	return s.r.GetProduct(ctx, id)
}

func (s *productService) CreateProduct(ctx context.Context, product *model.CreateProductRequest) (*model.Product, error) {
	if product.Name == "" {
		return nil, errors.New("Name" + ErrValidate)
	}
	if product.Price <= 0 {
		return nil, errors.New("Price" + ErrValidate)
	}
	if product.CategoryId == uuid.Nil {
		return nil, errors.New("CategoryId" + ErrValidate)
	}
	return s.r.CreateProduct(ctx, product)
}

func (s *productService) GetProductByFilter(ctx context.Context, filter *dto.PageRequest) (*dto.PageResponse, error) {
	if filter.Size <= 0 {
		filter.Size = 10
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}
	return s.r.GetProductByFilter(ctx, filter)
}

func (s *productService) UpdateProduct(ctx context.Context, id uuid.UUID, product *model.Product) (*model.Product, error) {
	if product.Name == "" {
		return nil, errors.New("Name" + ErrValidate)
	}
	if product.Price <= 0 {
		return nil, errors.New("Price" + ErrValidate)
	}
	return s.r.UpdateProduct(ctx, id, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("Id" + ErrValidate)
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("Id" + ErrValidate)
	}

	return s.r.DeleteProduct(ctx, idUuid)
}
