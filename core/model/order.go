package dto

type CreateOrderRequest struct {
	CustomerID uint  `json:"customer_id"`
	ProductID  uint  `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}

type UpdateOrderRequest struct {
	ID         uint  `uri:"id"`
	CustomerID uint  `json:"customer_id"`
	ProductID  uint  `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}

type OrderResponse struct {
	ID         uint  `json:"id"`
	CustomerID uint  `json:"customer_id"`
	ProductID  uint  `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}

type OrderRequest struct {
	ID uint `uri:"id"`
}
