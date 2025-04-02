package model

type CreateInventoryRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type UpdateInventoryRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}
