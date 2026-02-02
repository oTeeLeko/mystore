package repository

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	dto "github.com/oTeeLeko/mystore/core/model"
	tbl "github.com/oTeeLeko/mystore/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Check Inventory
		var inventory tbl.Inventory
		if err := tx.WithContext(ctx).Where("product_id = ?", req.ProductID).First(&inventory).Error; err != nil {
			return err
		}

		if inventory.Quantity == nil || int32(*inventory.Quantity) < req.Quantity {
			return errors.New("insufficient stock")
		}

		// 2. Deduct Inventory
		newQty := int(*inventory.Quantity) - int(req.Quantity)
		inventory.Quantity = &newQty
		if err := tx.WithContext(ctx).Save(&inventory).Error; err != nil {
			return err
		}

		// 3. Create Order
		var entity tbl.Orders
		copier.Copy(&entity, req)
		if err := tx.WithContext(ctx).Create(&entity).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, req *dto.UpdateOrderRequest) error {
	var entity tbl.Orders
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	copier.CopyWithOption(&entity, req, copier.Option{IgnoreEmpty: true})
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, req *dto.OrderRequest) error {
	var entity tbl.Orders
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&entity).Error
}

func (r *OrderRepository) GetOrder(ctx context.Context, req *dto.OrderRequest) (*dto.OrderResponse, error) {
	var entity tbl.Orders
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return nil, err
	}
	var response dto.OrderResponse
	copier.Copy(&response, &entity)
	return &response, nil
}

func (r *OrderRepository) GetListOrders(ctx context.Context) ([]*dto.OrderResponse, error) {
	var entities []tbl.Orders
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	var responses []*dto.OrderResponse
	copier.Copy(&responses, &entities)
	return responses, nil
}
