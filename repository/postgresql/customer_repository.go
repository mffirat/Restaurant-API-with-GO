package postgresql

import (
	"Go2/model"

	"gorm.io/gorm"
)


type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (r *CustomerRepo) CreateCustomer(c *model.Customer) error {
	return r.db.Create(&c).Error
}

func (r *CustomerRepo) UpdateCustomer(c model.Customer) error {
	return r.db.Save(&c).Error
}

func (r *CustomerRepo) GetCustomerByID(id uint) (model.Customer, error) {
	var c model.Customer
	err := r.db.First(&c, id).Error
	return c, err
}

func (r *CustomerRepo) GetTotalCustomers(start, end string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Customer{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Count(&count).Error
	return count, err
}

func (r *CustomerRepo) GetChildrenCount(start, end string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Customer{}).
		Where("age_group = ? AND created_at BETWEEN ? AND ?", "child", start, end).
		Count(&count).Error
	return count, err
}

func (r *CustomerRepo) GetTotalIncome(start, end string) (float64, error) {
	var total float64
	err := r.db.Model(&model.Customer{}).
		Where("exited_at BETWEEN ? AND ?", start, end).
		Select("SUM(payment)").Scan(&total).Error
	return total, err
}
