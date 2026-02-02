package dto

type CreateInventoryRequest struct {
	ProductID uint  `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type UpdateInventoryRequest struct {
	ID        uint  `uri:"id"`
	ProductID uint  `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type InventoryResponse struct {
	ID        uint  `json:"id"`
	ProductID uint  `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type InventoryRequest struct {
	ID uint `uri:"id"`
}
