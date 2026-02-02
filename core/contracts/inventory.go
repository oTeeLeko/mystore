package contracts

import (
	"context"

	dto "github.com/oTeeLeko/mystore/core/model"
)

type InventoryRepository interface {
	CreateInventory(ctx context.Context, req *dto.CreateInventoryRequest) error
	UpdateInventory(ctx context.Context, req *dto.UpdateInventoryRequest) error
	DeleteInventory(ctx context.Context, req *dto.InventoryRequest) error
	GetInventory(ctx context.Context, req *dto.InventoryRequest) (*dto.InventoryResponse, error)
	GetListInventories(ctx context.Context) ([]*dto.InventoryResponse, error)
}
