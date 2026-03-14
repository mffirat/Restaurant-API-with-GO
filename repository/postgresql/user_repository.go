package postgresql

import (
	"Go2/domain/user"

	"gorm.io/gorm"
	"context"
	"Go2/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) CreateUser(ctx context.Context, u *user.User) error {
	ctx, span := tracing.Tracer.Start(ctx, "postgres.CreateUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.username", u.Username))
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	ctx, span := tracing.Tracer.Start(ctx, "postgres.GetByUsername")
	defer span.End()

	span.SetAttributes(attribute.String("user.username", username))
	var u user.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return &u, nil
}
