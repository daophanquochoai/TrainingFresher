package repository

import (
	"database/sql"
	"go-db-demo/internal/model"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	GetProductById(id string) (*model.Product, error)
}

type PostgresProductRepository struct {
	db *sql.DB
}

// init a new PostgresProductRepository
func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
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
