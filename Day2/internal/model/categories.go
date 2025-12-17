package model

type Categories struct {
	Id   string
	Name string
}

type CategoryAndProduct struct {
	Categories
	Product []Product
}

type CategoryRequest struct {
	CategoryId string   `json:"category_id"`
	ProductIds []string `json:"product_ids"`
}
