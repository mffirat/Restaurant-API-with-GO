package domain

import (
	"testing"
	"Go2/domain/user"
	"Go2/model"
)

type fakeCustomerRepo struct{}

func (f *fakeCustomerRepo) CreateCustomer(c *model.Customer) error { return nil }
func (f *fakeCustomerRepo) UpdateCustomer(c model.Customer) error  { return nil }
func (f *fakeCustomerRepo) GetCustomerByID(id uint) (model.Customer, error) {return model.Customer{}, nil}

func (f *fakeCustomerRepo) GetTotalCustomers(start, end string) (int64, error) { return 0, nil }
func (f *fakeCustomerRepo) GetChildrenCount(start, end string) (int64, error)  { return 5, nil }
func (f *fakeCustomerRepo) GetTotalIncome(start, end string) (float64, error)  { return 0, nil }


type fakeFloorRepo struct{}

func (f *fakeFloorRepo) IncreaseFloorCount(floor int) error   { return nil }
func (f *fakeFloorRepo) DecreaseFloorCount(floor int) error   { return nil }
func (f *fakeFloorRepo) GetFloorCount(floor int) (int, error) { return 0, nil }



type fakeUserRepo struct{}

func (f *fakeUserRepo) CreateUser(u *user.User) error {return nil}
func (f *fakeUserRepo) GetByUsername(username string) (*user.User, error) {return nil, nil}



func TestGetChildrenCount(t *testing.T) {
	customerRepo := &fakeCustomerRepo{}
	floorRepo := &fakeFloorRepo{}
	userRepo := &fakeUserRepo{}
	service := NewDomainService(customerRepo, floorRepo,userRepo)

	total, err := service.GetChildrenCount("2025-11-11", "2025-11-12")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := int64(5)
	if total != expected {
		t.Errorf("Expected %v,got %v", expected, total)
	}

}
