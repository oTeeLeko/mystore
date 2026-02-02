package repository

import (
	"context"

	"github.com/jinzhu/copier"
	dto "github.com/oTeeLeko/mystore/core/model"
	tbl "github.com/oTeeLeko/mystore/models"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, req *dto.CreateInventoryRequest) error {
	var entity tbl.Inventory
	copier.Copy(&entity, req)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *InventoryRepository) UpdateInventory(ctx context.Context, req *dto.UpdateInventoryRequest) error {
	var entity tbl.Inventory
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	copier.CopyWithOption(&entity, req, copier.Option{IgnoreEmpty: true})
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *InventoryRepository) DeleteInventory(ctx context.Context, req *dto.InventoryRequest) error {
	var entity tbl.Inventory
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&entity).Error
}

func (r *InventoryRepository) GetInventory(ctx context.Context, req *dto.InventoryRequest) (*dto.InventoryResponse, error) {
	var entity tbl.Inventory
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return nil, err
	}
	var response dto.InventoryResponse
	copier.Copy(&response, &entity)
	return &response, nil
}

func (r *InventoryRepository) GetListInventories(ctx context.Context) ([]*dto.InventoryResponse, error) {
	var entities []tbl.Inventory
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	var responses []*dto.InventoryResponse
	copier.Copy(&responses, &entities)
	return responses, nil
}
