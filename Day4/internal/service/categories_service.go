package service

import (
	"context"
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type CategoryService interface {
	GetCategoryById(ctx context.Context, categoryId string) (*model.Categories, error)
	CreateCategory(ctx context.Context, category *model.Categories) (*model.Categories, error)
	UpdateCategoryById(ctx context.Context, categoryId string, category *model.Categories) (*model.Categories, error)
	DeleteCategoryById(ctx context.Context, categoryId string) error
	AddProductIntoCategory(ctx context.Context, categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

// init
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repository}
}

// get category by id
func (s *categoryService) GetCategoryById(ctx context.Context, categoryId string) (*model.Categories, error) {
	return s.repo.GetCategoryById(ctx, categoryId)
}

// create category
func (s *categoryService) CreateCategory(ctx context.Context, category *model.Categories) (*model.Categories, error) {
	return s.repo.CreateCategory(ctx, category)
}

// update category
func (s *categoryService) UpdateCategoryById(ctx context.Context, categoryId string, category *model.Categories) (*model.Categories, error) {
	return s.repo.UpdateCategoryById(ctx, categoryId, category)
}

// delete category
func (s *categoryService) DeleteCategoryById(ctx context.Context, categoryId string) error {
	return s.repo.DeleteCategoryById(ctx, categoryId)
}

func (s *categoryService) AddProductIntoCategory(ctx context.Context, categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error) {
	return s.repo.AddProductIntoCategory(ctx, categoryParam)
}
