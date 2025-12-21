package dto

type PageRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type PageResponse struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
