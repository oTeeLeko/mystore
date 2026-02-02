package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	dto "github.com/oTeeLeko/mystore/core/model"
	tbl "github.com/oTeeLeko/mystore/models"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) CreateCustomer(ctx context.Context, req *dto.CreateCustomerRequest) error {
	var entity tbl.Customers
	copier.Copy(&entity, req)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *CustomerRepository) UpdateCustomer(ctx context.Context, req *dto.UpdateCustomerRequest) error {
	var entity tbl.Customers
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	copier.CopyWithOption(&entity, req, copier.Option{IgnoreEmpty: true})
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *CustomerRepository) DeleteCustomer(ctx context.Context, req *dto.CustomerRequest) error {
	var entity tbl.Customers
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&entity).Error
}

func (r *CustomerRepository) GetCustomer(ctx context.Context, req *dto.CustomerRequest) (*dto.GetCustomerResponse, error) {
	var entity tbl.Customers
	if err := r.db.WithContext(ctx).First(&entity, req.ID).Error; err != nil {
		return nil, err
	}
	var response dto.GetCustomerResponse
	copier.Copy(&response, &entity)
	return &response, nil
}

func (r *CustomerRepository) GetListCustomers(ctx context.Context) ([]*dto.GetCustomerResponse, error) {
	var entities []tbl.Customers
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	var responses []*dto.GetCustomerResponse
	copier.Copy(&responses, &entities)
	return responses, nil
}
