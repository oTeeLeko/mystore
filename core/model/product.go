package dto

import "github.com/shopspring/decimal"

type CreateProductRequest struct {
	Name  *string          `json:"name"`
	Price *decimal.Decimal `json:"price"`
}

type UpdateProductRequest struct {
	ID    uint             `uri:"id"`
	Name  *string          `json:"name"`
	Price *decimal.Decimal `json:"price"`
}

type ProductResponse struct {
	ID    uint             `json:"id"`
	Name  *string          `json:"name"`
	Price *decimal.Decimal `json:"price"`
}

type ProductRequest struct {
	ID uint `uri:"id"`
}
