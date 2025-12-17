package service

import (
	"go-db-demo/internal/model"
	"go-db-demo/internal/repository"
)

type CategoryService interface {
	GetCategoryById(categoryId string) (*model.Categories, error)
	CreateCategory(category *model.Categories) (*model.Categories, error)
	UpdateCategoryById(categoryId string, category *model.Categories) (*model.Categories, error)
	DeleteCategoryById(categoryId string) error
	AddProductIntoCategory(categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error)
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

// update category
func (s *categoryService) UpdateCategoryById(categoryId string, category *model.Categories) (*model.Categories, error) {
	return s.repo.UpdateCategoryById(categoryId, category)
}

// delete category
func (s *categoryService) DeleteCategoryById(categoryId string) error {
	return s.repo.DeleteCategoryById(categoryId)
}

func (s *categoryService) AddProductIntoCategory(categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error) {
	return s.repo.AddProductIntoCategory(categoryParam)
}
