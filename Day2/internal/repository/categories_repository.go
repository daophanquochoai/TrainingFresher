package repository

import (
	"database/sql"
	"go-db-demo/internal/model"
)

type CategoryRepository interface {
	CreateCategory(category *model.Categories) (*model.Categories, error)
	GetCategoryById(categoryId string) (*model.Categories, error)
}

type PostgresCategoryRepository struct {
	db *sql.DB
}

// init a new PostgresCategoryRepository
func NewPostgresCategoryRepository(db *sql.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

// create a new category
func (r *PostgresCategoryRepository) CreateCategory(
	category *model.Categories,
) (*model.Categories, error) {

	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, name
	`
	row := r.db.QueryRow(query, category.Name)

	var created model.Categories
	if err := row.Scan(&created.Id, &created.Name); err != nil {
		return nil, err
	}

	return &created, nil
}

// get category by id
func (r *PostgresCategoryRepository) GetCategoryById(categoryId string) (*model.Categories, error) {
	quey := "SELECT id, name FROM categories WHERE id = $1"
	row := r.db.QueryRow(quey, categoryId)
	category := &model.Categories{}
	err := row.Scan(&category.Id, &category.Name)
	return category, err
}
