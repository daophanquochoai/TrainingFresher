package model

import (
	"category/pb/productservice/proto/productpb"
	"github.com/google/uuid"
)

type Category struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `json:"name" gorm:"unique"`
}

type CategoryProduct struct {
	Cate Category                      `json:"category"`
	Prod productpb.ProductListResponse `json:"product"`
}
