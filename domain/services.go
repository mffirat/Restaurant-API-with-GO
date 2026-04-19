package domain

import (
	"Go2/domain/user"
	"Go2/model"
	"fmt"
	"os"

	"time"

	"context"

	"github.com/golang-jwt/jwt/v5"
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

func (s *DomainService) EnterCustomer(ctx context.Context,tenantID uint, gender, ageGroup string, floor int) (*model.Customer, error) {

	customer := &model.Customer{
		TenantID:  tenantID,
		Gender:    gender,
		AgeGroup:  ageGroup,
		Floor:     floor,
		Payment:   0,
		EnteredAt: time.Now(),
	}
	if err := s.customerRepo.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}
	if err := s.floorRepo.IncreaseFloorCount(ctx, floor); err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *DomainService) ExitCustomer(ctx context.Context,tenantID uint, id uint, payment float64) error {

	customer, err := s.customerRepo.GetCustomerByID(ctx,tenantID , id)
	if err != nil {
		return err
	}
	now := time.Now()
	customer.Payment = payment
	customer.ExitedAt = &now

	if err := s.customerRepo.UpdateCustomer(ctx, customer); err != nil {
		return err
	}
	return s.floorRepo.DecreaseFloorCount(ctx, customer.Floor)
}
func (s *DomainService) RegisterUser(ctx context.Context,tenantID uint, username, password string) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &user.User{
		TenantID: tenantID,
		Username: username,
		Password: string(hashed),
		Role:     "user",
	}

	return s.userRepo.CreateUser(ctx, u)
}
func (s *DomainService) LoginUser(ctx context.Context, username, password string) (string, error) {

	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("Invalid Entrance")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("Invalid Entrance")
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"tenant_id": user.TenantID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", fmt.Errorf("JWT_SECRET not set")
	}
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *DomainService) GetCounts(ctx context.Context) (model.FloorCount, error) {

	f1, err := s.floorRepo.GetFloorCount(ctx, 1)
	if err != nil {
		return model.FloorCount{}, err
	}
	f2, err := s.floorRepo.GetFloorCount(ctx, 2)
	if err != nil {
		return model.FloorCount{}, err
	}
	f3, err := s.floorRepo.GetFloorCount(ctx, 3)
	if err != nil {
		return model.FloorCount{}, err
	}
	total := f1 + f2 + f3
	return model.FloorCount{Floor1: f1, Floor2: f2, Floor3: f3, Total: total}, nil
}

func (s *DomainService) GetTotalCustomers(ctx context.Context,tenantID uint, start, end string) (int64, error) {

	return s.customerRepo.GetTotalCustomers(ctx,tenantID , start, end)
}

func (s *DomainService) GetChildrenCount(ctx context.Context,tenantID uint, start, end string) (int64, error) {

	return s.customerRepo.GetChildrenCount(ctx,tenantID , start, end)
}

func (s *DomainService) GetTotalIncome(ctx context.Context,tenantID uint, start, end string) (float64, error) {

	return s.customerRepo.GetTotalIncome(ctx,tenantID , start, end)
}
