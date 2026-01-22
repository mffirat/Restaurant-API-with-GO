package domain

import (
	"Go2/domain/user"
	"Go2/model"
	"fmt"

	"time"

	"golang.org/x/crypto/bcrypt"
)

type DomainService struct {
	customerRepo CustomerRepoInterface
	floorRepo    FloorRepoInterface
	userRepo     UserRepoInterface
}

func NewDomainService(customerRepo CustomerRepoInterface, floorRepo FloorRepoInterface, userRepo UserRepoInterface) *DomainService {
	return &DomainService{
		customerRepo: customerRepo,
		floorRepo:    floorRepo,
		userRepo:     userRepo,
	}
}

func (s *DomainService) EnterCustomer(gender, ageGroup string, floor int) (*model.Customer, error) {
	customer := &model.Customer{
		Gender:    gender,
		AgeGroup:  ageGroup,
		Floor:     floor,
		Payment:   0,
		EnteredAt: time.Now(),
	}
	if err := s.customerRepo.CreateCustomer(customer); err != nil {
		return nil, err
	}
	if err := s.floorRepo.IncreaseFloorCount(floor); err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *DomainService) ExitCustomer(id uint, payment float64) error {
	customer, err := s.customerRepo.GetCustomerByID(id)
	if err != nil {
		return err
	}
	now := time.Now()
	customer.Payment = payment
	customer.ExitedAt = &now

	if err := s.customerRepo.UpdateCustomer(customer); err != nil {
		return err
	}
	return s.floorRepo.DecreaseFloorCount(customer.Floor)
}
func (s *DomainService) RegisterUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &user.User{
		Username: username,
		Password: string(hashed),
		Role:     "user",
	}

	return s.userRepo.CreateUser(u)
}
func (s *DomainService) LoginUser(username, password string) error {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return fmt.Errorf("Invalid Entrance")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return fmt.Errorf("Invalid Entrance")
	}
	return nil
}

func (s *DomainService) GetCounts() (model.FloorCount, error) {
	f1, err := s.floorRepo.GetFloorCount(1)
	if err != nil {
		return model.FloorCount{}, err
	}
	f2, err := s.floorRepo.GetFloorCount(2)
	if err != nil {
		return model.FloorCount{}, err
	}
	f3, err := s.floorRepo.GetFloorCount(3)
	if err != nil {
		return model.FloorCount{}, err
	}
	total := f1 + f2 + f3
	return model.FloorCount{Floor1: f1, Floor2: f2, Floor3: f3, Total: total}, nil
}

func (s *DomainService) GetTotalCustomers(start, end string) (int64, error) {
	return s.customerRepo.GetTotalCustomers(start, end)
}

func (s *DomainService) GetChildrenCount(start, end string) (int64, error) {
	return s.customerRepo.GetChildrenCount(start, end)
}

func (s *DomainService) GetTotalIncome(start, end string) (float64, error) {
	return s.customerRepo.GetTotalIncome(start, end)
}
