package contracts

import (
	"context"

	dto "github.com/oTeeLeko/mystore/core/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error
	UpdateOrder(ctx context.Context, req *dto.UpdateOrderRequest) error
	DeleteOrder(ctx context.Context, req *dto.OrderRequest) error
	GetOrder(ctx context.Context, req *dto.OrderRequest) (*dto.OrderResponse, error)
	GetListOrders(ctx context.Context) ([]*dto.OrderResponse, error)
}
