package model

type CreateOrderRequest struct {
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Quantity   int32  `json:"quantity"`
}
