package domain

import (
	"Go2/domain/user"
	"Go2/model"
	"fmt"
	"os"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"context"
	"Go2/tracing"
	"go.opentelemetry.io/otel/attribute"
	
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

func (s *DomainService) EnterCustomer(ctx context.Context,gender, ageGroup string, floor int) (*model.Customer, error) {
	ctx, span := tracing.StartSpan(ctx, "service.EnterCustomer")
	defer span.End()

	span.SetAttributes(
		attribute.String("customer.gender", gender),
		attribute.String("customer.age_group", ageGroup),
		attribute.Int("customer.floor", floor),
	)
	customer := &model.Customer{
		Gender:    gender,
		AgeGroup:  ageGroup,
		Floor:     floor,
		Payment:   0,
		EnteredAt: time.Now(),
	}
	if err := s.customerRepo.CreateCustomer(ctx,customer); err != nil {
		return nil, err
	}
	if err := s.floorRepo.IncreaseFloorCount(ctx,floor); err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *DomainService) ExitCustomer(ctx context.Context,id uint, payment float64) error {
	ctx, span := tracing.StartSpan(ctx, "service.ExitCustomer")
	defer span.End()

	span.SetAttributes(
		attribute.Int("customer.id", int(id)),
		attribute.Float64("customer.payment", payment),
	)
	customer, err := s.customerRepo.GetCustomerByID(ctx,id)
	if err != nil {
		return err
	}
	now := time.Now()
	customer.Payment = payment
	customer.ExitedAt = &now

	if err := s.customerRepo.UpdateCustomer(ctx,customer); err != nil {
		return err
	}
	return s.floorRepo.DecreaseFloorCount(ctx,customer.Floor)
}
func (s *DomainService) RegisterUser(ctx context.Context,username, password string) error {
	ctx, span := tracing.StartSpan(ctx, "service.RegisterUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.username", username))
	
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &user.User{
		Username: username,
		Password: string(hashed),
		Role:     "user",
	}

	return s.userRepo.CreateUser(ctx,u)
}
func (s *DomainService) LoginUser(ctx context.Context,username, password string) (string, error) {
	ctx, span := tracing.StartSpan(ctx, "service.LoginUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.username", username))

	user, err := s.userRepo.GetByUsername(ctx,username)
	if err != nil {
		return "",fmt.Errorf("Invalid Entrance")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "",fmt.Errorf("Invalid Entrance")
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret:= []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
	return "", fmt.Errorf("JWT_SECRET not set")
}
	signedToken,err :=token.SignedString(secret)
	if err !=nil{
		return "",err
	}
	return signedToken,nil
}

func (s *DomainService) GetCounts(ctx context.Context) (model.FloorCount, error) {
	ctx, span := tracing.StartSpan(ctx, "service.GetCounts")
	defer span.End()
	f1, err := s.floorRepo.GetFloorCount(ctx,1)
	if err != nil {
		return model.FloorCount{}, err
	}
	f2, err := s.floorRepo.GetFloorCount(ctx,2)
	if err != nil {
		return model.FloorCount{}, err
	}
	f3, err := s.floorRepo.GetFloorCount(ctx,3)
	if err != nil {
		return model.FloorCount{}, err
	}
	total := f1 + f2 + f3
	return model.FloorCount{Floor1: f1, Floor2: f2, Floor3: f3, Total: total}, nil
}

func (s *DomainService) GetTotalCustomers(ctx context.Context,start, end string) (int64, error) {
	ctx, span := tracing.StartSpan(ctx, "service.GetTotalCustomers")
	defer span.End()
	return s.customerRepo.GetTotalCustomers(ctx,start, end)
}

func (s *DomainService) GetChildrenCount(ctx context.Context,start, end string) (int64, error) {
	ctx, span := tracing.StartSpan(ctx, "service.GetChildrenCount")
	defer span.End()
	return s.customerRepo.GetChildrenCount(ctx,start, end)
}

func (s *DomainService) GetTotalIncome(ctx context.Context,start, end string) (float64, error) {
	ctx, span := tracing.StartSpan(ctx, "service.GetTotalIncome")
	defer span.End()
	return s.customerRepo.GetTotalIncome(ctx,start, end)
}
