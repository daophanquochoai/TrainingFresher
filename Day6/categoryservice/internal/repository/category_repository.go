package repository

import (
	"category/internal/dto"
	"category/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

var baseCategory = "category:"
var ErrCategoryNotFound = errors.New("Category not found")

type CategoryRepository interface {
	GetCategory(ctx context.Context, id uuid.UUID) (*model.Category, error)
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, id string, category *model.Category) (*model.Category, error)
	GetCategoryList(ctx context.Context, pageRequest *dto.PageRequest) (*dto.PageResponse, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
	r  *redis.Client
}

func NewCategoryRepository(db *gorm.DB, r *redis.Client) CategoryRepository {
	return &categoryRepository{
		db: db,
		r:  r,
	}
}

func (repo *categoryRepository) GetCategoryList(ctx context.Context, pageRequest *dto.PageRequest) (*dto.PageResponse, error) {
	var categories []*model.Category
	var total int64

	offset := (pageRequest.Page - 1) * pageRequest.Page

	query := repo.db.WithContext(ctx).Model(&model.Category{})

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := query.Limit(pageRequest.Size).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}

	return &dto.PageResponse{
		Page:  pageRequest.Page,
		Size:  pageRequest.Size,
		Total: total,
		Data:  categories,
	}, nil
}

func (repo *categoryRepository) GetCategory(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	// redis
	key := baseCategory + id.String()
	cached, err := repo.r.Get(ctx, key).Bytes()
	if err == nil && cached != nil {
		var category model.Category
		if err := json.Unmarshal(cached, &category); err == nil {
			fmt.Println("Using data of redis")
			return nil, err
		}
	}

	// query
	var p model.Category
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	_ = repo.r.Set(ctx, key, &p, 5*time.Minute)
	return &p, nil
}

func (repo *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {

	tx := repo.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result model.Category
	queryCategory := "INSERT INTO categories (name) VALUES (?) RETURNING id, name"
	if err := tx.Raw(queryCategory, category.Name).Scan(&result).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &result, nil
}

func (repo *categoryRepository) UpdateCategory(ctx context.Context, id string, category *model.Category) (*model.Category, error) {

	var result model.Category
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, ErrCategoryNotFound
	}

	updated := repo.db.WithContext(ctx).Where("id = ?", id).Updates(category).Model(&result)
	if err := updated.Error; err != nil {
		return nil, err
	}

	var updatedCategory model.Category
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&updatedCategory).Error; err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}

func (repo *categoryRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	result := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Category{})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return ErrCategoryNotFound
	}
	return nil
}
