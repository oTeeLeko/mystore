package contracts

import (
	"context"

	dto "github.com/oTeeLeko/mystore/core/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error
	UpdateProduct(ctx context.Context, req *dto.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, req *dto.ProductRequest) error
	GetProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error)
	GetListProducts(ctx context.Context) ([]*dto.ProductResponse, error)
}
