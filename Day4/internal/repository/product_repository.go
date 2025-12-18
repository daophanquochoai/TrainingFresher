package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/model"
	"time"
)

var baseKeyProduct string = "product:"

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	DeleteProductById(ctx context.Context, id string) error
	UpdateProductById(ctx context.Context, id string, product *model.Product) (*model.Product, error)
}

type PostgresProductRepository struct {
	db    *sql.DB
	redis *redis.Client
}

// init a new PostgresProductRepository
func NewPostgresProductRepository(db *sql.DB, redis *redis.Client) ProductRepository {
	return &PostgresProductRepository{
		db:    db,
		redis: redis,
	}
}

// create a product
func (r *PostgresProductRepository) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	query := `INSERT INTO products ( name, price) VALUES ($1, $2) RETURNING id, name, price`
	row := r.db.QueryRowContext(ctx, query, product.Name, product.Price)

	var createProduct model.Product
	if err := row.Scan(&createProduct.Id, &createProduct.Name, &createProduct.Price); err != nil {
		return nil, err
	}
	return &createProduct, nil
}

// get product by id
func (r *PostgresProductRepository) GetProductById(ctx context.Context, id string) (*model.Product, error) {

	// redis
	key := baseKeyProduct + id
	cached, errCache := r.redis.Get(ctx, key).Bytes()
	if errCache == nil {
		var product model.Product
		if errUnmarshal := json.Unmarshal(cached, &product); errUnmarshal == nil {
			fmt.Println("INFO : Using product in redis")
			return &product, nil
		}
	}

	// query
	query := "SELECT id, name, price FROM products WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	product := &model.Product{}
	// error
	err := row.Scan(&product.Id, &product.Name, &product.Price)

	// cache
	b, _ := json.Marshal(product)
	_ = r.redis.Set(ctx, key, b, 60*time.Second).Err()

	// return
	return product, err
}

// delete product by id
func (r *PostgresProductRepository) DeleteProductById(ctx context.Context, id string) error {
	// query
	query := "SELECT delete_product_by_id($1)"
	_, err := r.db.ExecContext(ctx, query, id)

	// delete cache
	fmt.Println("Delete product in redis")
	key := baseKey + id
	_ = r.redis.Del(ctx, key).Err()

	return err
}

// update product by id
func (r *PostgresProductRepository) UpdateProductById(ctx context.Context, id string, product *model.Product) (*model.Product, error) {
	// query
	query := "SELECT id, name, price FROM update_product_by_id($1,$2,$3)"
	row := r.db.QueryRowContext(ctx, query, id, product.Name, product.Price)
	// mapping data
	var productUpdated model.Product
	if err := row.Scan(&productUpdated.Id, &productUpdated.Name, &productUpdated.Price); err != nil {
		return nil, err
	}

	// cache
	fmt.Println("Delete into in redis")
	key := baseKey + id
	_ = r.redis.Del(ctx, key).Err()
	return &productUpdated, nil
}
