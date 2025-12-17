package repository

import (
	"database/sql"
	"go-db-demo/internal/model"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	GetProductById(id string) (*model.Product, error)
	DeleteProductById(id string) error
	UpdateProductById(id string, product *model.Product) (*model.Product, error)
}

type PostgresProductRepository struct {
	db *sql.DB
}

// init a new PostgresProductRepository
func NewPostgresProductRepository(db *sql.DB) ProductRepository {
	return &PostgresProductRepository{db: db}
}

// create a product
func (r *PostgresProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	query := `INSERT INTO products ( name, price) VALUES ($1, $2) RETURNING id, name, price`
	row := r.db.QueryRow(query, product.Name, product.Price)
	var createProduct model.Product
	if err := row.Scan(&createProduct.Id, &createProduct.Name, &createProduct.Price); err != nil {
		return nil, err
	}
	return &createProduct, nil
}

// get product by id
func (r *PostgresProductRepository) GetProductById(id string) (*model.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1"
	row := r.db.QueryRow(query, id)
	product := &model.Product{}
	err := row.Scan(&product.Id, &product.Name, &product.Price)
	return product, err
}

// delete product by id
func (r *PostgresProductRepository) DeleteProductById(id string) error {
	query := "SELECT delete_product_by_id($1)"
	_, err := r.db.Exec(query, id)
	return err
}

// update product by id
func (r *PostgresProductRepository) UpdateProductById(id string, product *model.Product) (*model.Product, error) {
	query := "SELECT id, name, price FROM update_product_by_id($1,$2,$3)"
	row := r.db.QueryRow(query, id, product.Name, product.Price)
	var productUpdated model.Product
	if err := row.Scan(&productUpdated.Id, &productUpdated.Name, &productUpdated.Price); err != nil {
		return nil, err
	}
	return &productUpdated, nil
}
