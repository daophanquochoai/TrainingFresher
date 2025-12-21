package model

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}

type CreateProductRequest struct {
	Name       string    `json:"name"`
	Price      float32   `json:"price"`
	CategoryId uuid.UUID `json:"categoryId"`
}
