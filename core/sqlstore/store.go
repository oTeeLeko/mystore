package store

import (
	"github.com/oTeeLeko/mystore/core/contracts"
	"github.com/oTeeLeko/mystore/core/repository"
	"gorm.io/gorm"
)

type Store struct {
	CustomerRepo  contracts.CustomerRepository
	ProductRepo   contracts.ProductRepository
	InventoryRepo contracts.InventoryRepository
	OrderRepo     contracts.OrderRepository
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		CustomerRepo:  repository.NewCustomerRepository(db),
		ProductRepo:   repository.NewProductRepository(db),
		InventoryRepo: repository.NewInventoryRepository(db),
		OrderRepo:     repository.NewOrderRepository(db),
	}
}
