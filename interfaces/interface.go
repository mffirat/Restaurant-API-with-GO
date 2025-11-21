package interfaces

import "Go2/model"

type CustomerRepoInterface interface {
	CreateCustomer(c *model.Customer) error
	UpdateCustomer(c model.Customer) error
	GetCustomerByID(id uint) (model.Customer, error)
	GetTotalCustomers(start, end string) (int64, error)
	GetChildrenCount(start, end string) (int64, error)
	GetTotalIncome(start, end string) (float64, error)
}
type FloorRepoInterface interface {
	IncreaseFloorCount(floor int) error
	DecreaseFloorCount(floor int) error
	GetFloorCount(floor int) (int, error)
}