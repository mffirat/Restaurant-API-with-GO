package postgresql

import (
	"Go2/model"

	"context"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, c *model.Customer) error {
	return r.db.Create(c).Error
}

func (r *CustomerRepo) UpdateCustomer(ctx context.Context, c model.Customer) error {
	return r.db.Save(&c).Error
}

func (r *CustomerRepo) GetCustomerByID(ctx context.Context, tenantID uint, id uint) (model.Customer, error) {
    var c model.Customer
    err := r.db.WithContext(ctx).
        Where("tenant_id = ? AND id = ?", tenantID, id).  
        First(&c).Error
    return c, err
}

func (r *CustomerRepo) GetTotalCustomers(ctx context.Context,tenantID uint, start, end string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("tenant_id = ? AND created_at BETWEEN ? AND ?",tenantID , start, end).
		Count(&count).Error
	return count, err
}

func (r *CustomerRepo) GetChildrenCount(ctx context.Context,tenantID uint, start, end string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("tenant_id = ? AND age_group = ? AND created_at BETWEEN ? AND ?",tenantID , "child", start, end).
		Count(&count).Error
	return count, err
}

func (r *CustomerRepo) GetTotalIncome(ctx context.Context,tenantID uint, start, end string) (float64, error) {
	var total float64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("tenant_id = ? AND exited_at BETWEEN ? AND ?",tenantID , start, end).
		Select("SUM(payment)").Scan(&total).Error
	return total, err
}
