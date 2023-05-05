package dto

import "time"

type NewProductRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty" example:"Rokok Surya 16"`
	Description string `json:"description" valid:"required~description cannot be empty" example:"Deskripsi dari Rokok Surya 16"`
}

type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ProductResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetProductsResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       []ProductResponse `json:"data"`
}
