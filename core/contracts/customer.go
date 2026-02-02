package contracts

import (
	"context"

	dto "github.com/oTeeLeko/mystore/core/model"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, req *dto.CreateCustomerRequest) error
	UpdateCustomer(ctx context.Context, req *dto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, req *dto.CustomerRequest) error
	GetCustomer(ctx context.Context, req *dto.CustomerRequest) (*dto.GetCustomerResponse, error)
	GetListCustomers(ctx context.Context) ([]*dto.GetCustomerResponse, error)
}
