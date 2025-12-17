package repository

import (
	"database/sql"
	"go-db-demo/internal/model"
)

type CategoryRepository interface {
	CreateCategory(category *model.Categories) (*model.Categories, error)
	GetCategoryById(categoryId string) (*model.Categories, error)
	DeleteCategoryById(categoryId string) error
	UpdateCategoryById(categoryId string, category *model.Categories) (*model.Categories, error)
	AddProductIntoCategory(categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error)
}

type PostgresCategoryRepository struct {
	db *sql.DB
}

// init a new PostgresCategoryRepository
func NewPostgresCategoryRepository(db *sql.DB) CategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

// create a new category
func (r *PostgresCategoryRepository) CreateCategory(category *model.Categories) (*model.Categories, error) {

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

// update category by id
func (r *PostgresCategoryRepository) UpdateCategoryById(categoryId string, category *model.Categories) (*model.Categories, error) {
	query := "SELECT id, name FROM update_category_by_id($1, $2)"
	row := r.db.QueryRow(query, categoryId, category.Name)
	var categoryReturn model.Categories
	if err := row.Scan(&categoryReturn.Id, &categoryReturn.Name); err != nil {
		return nil, err
	}
	return &categoryReturn, nil
}

// delete category by id
func (r *PostgresCategoryRepository) DeleteCategoryById(categoryId string) error {
	query := "SELECT delete_category_by_id($1)"
	_, err := r.db.Exec(query, categoryId)
	return err
}

// add product into category
func (r *PostgresCategoryRepository) AddProductIntoCategory(categoryParam *model.CategoryRequest) (*model.CategoryAndProduct, error) {
	query := "SELECT id, name FROM categories c WHERE c.id = $1"
	rowCate := r.db.QueryRow(query, categoryParam.CategoryId)
	var categorySel model.Categories
	if err := rowCate.Scan(&categorySel.Id, &categorySel.Name); err != nil {
		return nil, err
	}
	// get product have
	queryCatePro := "SELECT pc.product_id FROM product_categories as pc WHERE pc.category_id = $1"
	rowCatePro, err := r.db.Query(queryCatePro, categoryParam.CategoryId)

	if err != nil {
		return nil, err
	}
	defer rowCatePro.Close()

	var existingProduct = make(map[string]bool)
	for rowCatePro.Next() {
		var pid string
		if er := rowCatePro.Scan(&pid); er != nil {
			return nil, er
		}
		existingProduct[pid] = true
	}
	// product from body
	var incomingProduct = make(map[string]bool)
	for _, pid := range categoryParam.ProductIds {
		incomingProduct[pid] = true
	}

	// product will add
	var toAdd []string
	for pid := range incomingProduct {
		if !existingProduct[pid] {
			toAdd = append(toAdd, pid)
		}
	}
	// product will remove
	var toRemove []string
	for pid := range existingProduct {
		if !incomingProduct[pid] {
			toRemove = append(toRemove, pid)
		}
	}
	// init transaction
	tx, err := r.db.Begin()

	if err != nil {
		return nil, err
	}
	// insert product need add
	insertQuery := "INSERT INTO product_categories(product_id,category_id) VALUES ($1, $2) ON CONFLICT DO NOTHING"

	for _, pid := range toAdd {
		print("pid :", pid)
		if _, err := tx.Exec(insertQuery, pid, categoryParam.CategoryId); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// delete product need remove
	deleteQuery := "DELETE FROM product_categories as pc WHERE pc.product_id = $1 AND pc.category_id = $2"

	for _, pid := range toRemove {
		if _, err := tx.Exec(deleteQuery, pid, categoryParam.CategoryId); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	// query product to return
	queryProductCategory := "SELECT p.id, p.name, p.price FROM product_categories as pc JOIN products p on p.id = pc.product_id WHERE pc.category_id = $1"
	rowProCate, err := r.db.Query(queryProductCategory, categoryParam.CategoryId)
	if err != nil {
		return nil, err
	}

	defer rowProCate.Close()
	var productSel []model.Product
	for rowProCate.Next() {
		var product model.Product
		if err := rowProCate.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		productSel = append(productSel, product)
	}

	categoryProduct := model.CategoryAndProduct{
		Categories: categorySel,
		Product:    productSel,
	}

	return &categoryProduct, nil
}
