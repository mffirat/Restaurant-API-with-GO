package domain

import (
	"Go2/domain/user"
	"Go2/model"
	"context"
)

type CustomerRepoInterface interface {
	CreateCustomer(ctx context.Context,c *model.Customer) error
	UpdateCustomer(ctx context.Context,c model.Customer) error
	GetCustomerByID(ctx context.Context,id uint) (model.Customer, error)
	GetTotalCustomers(ctx context.Context,start, end string) (int64, error)
	GetChildrenCount(ctx context.Context,start, end string) (int64, error)
	GetTotalIncome(ctx context.Context,start, end string) (float64, error)
}
type FloorRepoInterface interface {
	IncreaseFloorCount(ctx context.Context,floor int) error
	DecreaseFloorCount(ctx context.Context,floor int) error
	GetFloorCount(ctx context.Context,floor int) (int, error)
}
type UserRepoInterface interface {
	CreateUser(ctx context.Context,u *user.User) error
	GetByUsername(ctx context.Context,username string) (*user.User, error)
}


