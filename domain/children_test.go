package domain

import (
	"testing"
	"Go2/domain/user"
	"Go2/model"
	"context"
)

type fakeCustomerRepo struct{}

func (f *fakeCustomerRepo) CreateCustomer(ctx context.Context,c *model.Customer) error { return nil }
func (f *fakeCustomerRepo) UpdateCustomer(ctx context.Context,c model.Customer) error  { return nil }
func (f *fakeCustomerRepo) GetCustomerByID(ctx context.Context,tenantID uint,id uint) (model.Customer, error) {return model.Customer{}, nil}

func (f *fakeCustomerRepo) GetTotalCustomers(ctx context.Context,tenantID uint,start, end string) (int64, error) { return 0, nil }
func (f *fakeCustomerRepo) GetChildrenCount(ctx context.Context,tenantID uint,start, end string) (int64, error)  { return 5, nil }
func (f *fakeCustomerRepo) GetTotalIncome(ctx context.Context,tenantID uint,start, end string) (float64, error)  { return 0, nil }


type fakeFloorRepo struct{}

func (f *fakeFloorRepo) IncreaseFloorCount(ctx context.Context,floor int) error   { return nil }
func (f *fakeFloorRepo) DecreaseFloorCount(ctx context.Context,floor int) error   { return nil }
func (f *fakeFloorRepo) GetFloorCount(ctx context.Context,floor int) (int, error) { return 0, nil }



type fakeUserRepo struct{}

func (f *fakeUserRepo) CreateUser(ctx context.Context,u *user.User) error {return nil}
func (f *fakeUserRepo) GetByUsername(ctx context.Context,username string) (*user.User, error) {return nil, nil}



func TestGetChildrenCount(t *testing.T) {
	customerRepo := &fakeCustomerRepo{}
	floorRepo := &fakeFloorRepo{}
	userRepo := &fakeUserRepo{}
	service := NewDomainService(customerRepo, floorRepo,userRepo)

	ctx:=context.Background()
	total, err := service.GetChildrenCount(ctx,1,"2025-11-11", "2025-11-12")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := int64(5)
	if total != expected {
		t.Errorf("Expected %v,got %v", expected, total)
	}

}
