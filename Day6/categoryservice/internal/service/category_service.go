package service

import (
	"category/internal/dto"
	"category/internal/grpc/client"
	"category/internal/model"
	"category/internal/repository"
	"category/pb/productservice/proto/productpb"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
)

var ErrInvalid = "Is Invalid"
var ErrCategoryNameNotFound = "Category Name is empty"

type CategoryService interface {
	GetCategory(ctx context.Context, id string) (*model.CategoryProduct, error)
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, id string, category *model.Category) (*model.Category, error)
	GetCategoryList(ctx context.Context, pageRequest *dto.PageRequest) (*dto.PageResponse, error)
	DeleteCategory(ctx context.Context, id string) error
}

type categoryService struct {
	repo          repository.CategoryRepository
	productClient *client.ProductClient
}

func NewCategoryService(repo repository.CategoryRepository, product *client.ProductClient) CategoryService {
	return &categoryService{
		repo:          repo,
		productClient: product,
	}
}

func (s *categoryService) GetCategory(ctx context.Context, id string) (*model.CategoryProduct, error) {
	// Validate UUID
	idUuid, err := uuid.Parse(id)
	if id == "" || err != nil {
		return nil, fmt.Errorf("invalid category id: %s", id)
	}

	// Lấy category từ database
	category, err := s.repo.GetCategory(ctx, idUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	// Gọi gRPC để lấy products
	products, err := s.productClient.GetProductByCategoryId(ctx, id)
	if err != nil {
		log.Printf("[WARN] Failed to get products for category %s: %v", id, err)
		// Trả về category với products rỗng
		return &model.CategoryProduct{
			Cate: *category,
			Prod: productpb.ProductListResponse{}, // Empty response
		}, nil
	}

	// Log số lượng products (nếu cần debug)
	if products != nil && len(products.Products) > 0 {
		log.Printf("[INFO] Found %d products for category %s", len(products.Products), id)
	}

	return &model.CategoryProduct{
		Cate: *category,
		Prod: *products,
	}, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	if category.Name == "" {
		return nil, errors.New(ErrCategoryNameNotFound)
	}

	return s.repo.CreateCategory(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id string, category *model.Category) (*model.Category, error) {
	if category.Name == "" {
		return nil, errors.New(ErrCategoryNameNotFound)
	}

	return s.repo.UpdateCategory(ctx, id, category)
}

func (s *categoryService) GetCategoryList(ctx context.Context, pageRequest *dto.PageRequest) (*dto.PageResponse, error) {

	if pageRequest == nil {
		return nil, errors.New(ErrInvalid)
	}

	if pageRequest.Page <= 0 {
		pageRequest.Page = 1
	}

	if pageRequest.Size <= 0 {
		pageRequest.Size = 10
	}

	return s.repo.GetCategoryList(ctx, pageRequest)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id string) error {
	idUuid, err := uuid.Parse(id)
	if id == "" || err != nil {
		return errors.New("Id " + ErrInvalid)
	}
	return s.repo.DeleteCategory(ctx, idUuid)
}
