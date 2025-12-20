package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"product/internal/model"
	"time"
)

var baseProduct string = "product"

type ProductRepository interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
	rd *redis.Client
}

func NewProductRepository(db *gorm.DB, rd *redis.Client) ProductRepository {
	return &productRepository{
		db: db,
		rd: rd,
	}
}

func (r *productRepository) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	// redis
	key := baseProduct + id
	cached, err := r.rd.Get(ctx, key).Bytes()
	if err == nil && cached != nil {
		var product model.Product
		if e := json.Unmarshal(cached, &product); e != nil {
			return &product, nil
		}
	}

	// query
	var p model.Product
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}

	_ = r.rd.Set(ctx, key, p, 10*time.Minute)
	return &p, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	// begin transaction
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	// insert and select
	var result model.Product
	queryInsert := "INSERT INTO products (name, price) VALUES (?,?) RETURNING id, name, price"
	if err := tx.Raw(queryInsert, product.Name, product.Price).Scan(&result).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &result, nil
}
