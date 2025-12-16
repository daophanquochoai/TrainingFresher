package service

import (
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type CategoryService interface {
	GetCategoryById(categoryId string) (*model.Categories, error)
	CreateCategory(category *model.Categories) (*model.Categories, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

// init
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repository}
}

// get category by id
func (s *categoryService) GetCategoryById(categoryId string) (*model.Categories, error) {
	return s.repo.GetCategoryById(categoryId)
}

// create category
func (s *categoryService) CreateCategory(category *model.Categories) (*model.Categories, error) {
	return s.repo.CreateCategory(category)
}
