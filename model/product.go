package model

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductRequest struct {
	ID    string  `uri:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
