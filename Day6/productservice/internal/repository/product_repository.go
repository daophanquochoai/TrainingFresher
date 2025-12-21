package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"product/internal/dto"
	"product/internal/model"
	"time"
)

var baseProduct = "product:"
var baseProductCategory = "category_product:"
var ErrProductNotFound = "Product not found"

type ProductRepository interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.CreateProductRequest) (*model.Product, error)
	GetProductByFilter(ctx context.Context, filter *dto.PageRequest) (*dto.PageResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	GetProductByCategoryId(ctx context.Context, categoryId uuid.UUID) ([]*model.Product, error)
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

func (r *productRepository) GetProductByCategoryId(ctx context.Context, categoryId uuid.UUID) ([]*model.Product, error) {
	key := baseProductCategory + categoryId.String()
	cached, err := r.rd.Get(ctx, key).Bytes()
	if err == nil && cached != nil {
		var products []*model.Product
		err = json.Unmarshal(cached, &products)
		if err == nil {
			return products, nil
		}
	}

	var products []*model.Product
	if err := r.db.Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) > 0 {
		_ = r.rd.Set(ctx, key, products, 10*time.Minute)
	}

	return products, nil
}

func (r *productRepository) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	// redis
	key := baseProduct + id
	cached, err := r.rd.Get(ctx, key).Bytes()
	if err == nil && cached != nil {
		var product model.Product
		if e := json.Unmarshal(cached, &product); e == nil {
			return &product, nil
		}
	}

	// query
	var p model.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	_ = r.rd.Set(ctx, key, p, 10*time.Minute)
	return &p, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, product *model.CreateProductRequest) (*model.Product, error) {
	// begin transaction
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	// insert and select
	var result model.Product
	queryInsert := "INSERT INTO products (name, price, category_id) VALUES (?,?,?) RETURNING id, name, price"
	if err := tx.Raw(queryInsert, product.Name, product.Price, product.CategoryId).Scan(&result).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// delete redis
	key := baseProductCategory + product.CategoryId.String()
	_ = r.rd.Del(ctx, key).Err()

	return &result, nil
}

func (r *productRepository) GetProductByFilter(ctx context.Context, filter *dto.PageRequest) (*dto.PageResponse, error) {
	var products []*model.Product
	var total int64

	// Tính offset từ page và size
	offset := (filter.Page - 1) * filter.Size

	// Query với pagination
	query := r.db.WithContext(ctx).Model(&model.Product{})

	// Đếm tổng số records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Lấy data với limit và offset
	if err := query.Limit(filter.Size).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	// Trả về PageResponse
	return &dto.PageResponse{
		Page:  filter.Page,
		Size:  filter.Size,
		Total: int(total),
		Data:  products,
	}, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, id uuid.UUID, product *model.Product) (*model.Product, error) {

	var result model.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).Updates(product).Scan(&result).Error; err != nil {
		return nil, errors.New(ErrProductNotFound)
	}

	updated := r.db.WithContext(ctx).Where("id = ?", id).Updates(product).Model(&result)
	if updated.Error != nil {
		return nil, updated.Error
	}

	var updatedProduct model.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&updatedProduct).Error; err != nil {
		return nil, err
	}

	key := baseProduct + id.String()
	_ = r.rd.Del(ctx, key).Err()

	return &updatedProduct, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(ErrProductNotFound)
	}
	return nil
}
