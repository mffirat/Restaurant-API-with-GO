package postgresql

import (
	"Go2/model"

	"gorm.io/gorm"
	"Go2/tracing"
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)


type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, c *model.Customer) error {
	ctx, span := tracing.StartSpan(ctx, "postgres.CreateCustomer")
	defer span.End()

	if err := r.db.WithContext(ctx).Create(c).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	span.SetAttributes(attribute.Int("customer.id", int(c.ID)))
	return nil
}

func (r *CustomerRepo) UpdateCustomer(ctx context.Context, c model.Customer) error {
	ctx, span := tracing.StartSpan(ctx, "postgres.UpdateCustomer")
	defer span.End()

	span.SetAttributes(attribute.Int("customer.id", int(c.ID)))
	if err := r.db.WithContext(ctx).Save(&c).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func (r *CustomerRepo) GetCustomerByID(ctx context.Context, id uint) (model.Customer, error) {
	ctx, span := tracing.StartSpan(ctx, "postgres.GetCustomerByID")
	defer span.End()

	span.SetAttributes(attribute.Int("customer.id", int(id)))
	var c model.Customer
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c, err
	}
	return c, nil
}

func (r *CustomerRepo) GetTotalCustomers(ctx context.Context, start, end string) (int64, error) {
	ctx, span := tracing.StartSpan(ctx, "postgres.GetTotalCustomers")
	defer span.End()

	span.SetAttributes(
		attribute.String("query.start", start),
		attribute.String("query.end", end),
	)
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Count(&count).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return count, err
}

func (r *CustomerRepo) GetChildrenCount(ctx context.Context, start, end string) (int64, error) {
	ctx, span := tracing.StartSpan(ctx, "postgres.GetChildrenCount")
	defer span.End()

	var count int64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("age_group = ? AND created_at BETWEEN ? AND ?", "child", start, end).
		Count(&count).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return count, err
}

func (r *CustomerRepo) GetTotalIncome(ctx context.Context, start, end string) (float64, error) {
	ctx, span := tracing.StartSpan(ctx, "postgres.GetTotalIncome")
	defer span.End()

	var total float64
	err := r.db.WithContext(ctx).Model(&model.Customer{}).
		Where("exited_at BETWEEN ? AND ?", start, end).
		Select("SUM(payment)").Scan(&total).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return total, err
}
