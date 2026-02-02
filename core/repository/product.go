package repository

import (
	"context"

	"github.com/jinzhu/copier"
	dto "github.com/oTeeLeko/mystore/core/model"
	tbl "github.com/oTeeLeko/mystore/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error {
	var entity tbl.Products
	copier.Copy(&entity, req)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, req *dto.UpdateProductRequest) error {
	var entity tbl.Products
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	copier.CopyWithOption(&entity, req, copier.Option{
		IgnoreEmpty: true,
	})
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, req *dto.ProductRequest) error {
	var entity tbl.Products
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&entity).Error
}

func (r *ProductRepository) GetProduct(ctx context.Context, req *dto.ProductRequest) (*dto.ProductResponse, error) {
	var entity tbl.Products
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return nil, err
	}
	var response dto.ProductResponse
	copier.Copy(&response, &entity)
	return &response, nil
}

func (r *ProductRepository) GetListProducts(ctx context.Context) ([]*dto.ProductResponse, error) {
	var entities []tbl.Products
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	var responses []*dto.ProductResponse
	copier.Copy(&responses, &entities)
	return responses, nil
}
