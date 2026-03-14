package redis

import (
	"context"
	"fmt"

	"Go2/tracing"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type FloorRepo struct {
	client *redis.Client
}

func NewFloorRepo(client *redis.Client) *FloorRepo {
	return &FloorRepo{
		client: client,
	}
}

func (r *FloorRepo) IncreaseFloorCount(ctx context.Context, floor int) error {
	ctx, span := tracing.Tracer.Start(ctx, "redis.IncreaseFloorCount")
	defer span.End()

	key := fmt.Sprintf("floor:%d", floor)
	span.SetAttributes(attribute.String("redis.key", key))

	if err := r.client.Incr(ctx, key).Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func (r *FloorRepo) DecreaseFloorCount(ctx context.Context, floor int) error {
	ctx, span := tracing.Tracer.Start(ctx, "redis.DecreaseFloorCount")
	defer span.End()

	key := fmt.Sprintf("floor:%d", floor)
	span.SetAttributes(attribute.String("redis.key", key))

	if err := r.client.Decr(ctx, key).Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func (r *FloorRepo) GetFloorCount(ctx context.Context, floor int) (int, error) {
	ctx, span := tracing.Tracer.Start(ctx, "redis.GetFloorCount")
	defer span.End()

	key := fmt.Sprintf("floor:%d", floor)
	span.SetAttributes(attribute.String("redis.key", key))

	val, err := r.client.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return val, err
}

func (r *FloorRepo) InitializeFloorCounts(ctx context.Context, totalFloors int) error {
	ctx, span := tracing.Tracer.Start(ctx, "redis.InitializeFloorCounts")
	defer span.End()
	span.SetAttributes(attribute.Int("floors.total", totalFloors))

	for i := 1; i <= totalFloors; i++ {
		key := fmt.Sprintf("floor:%d", i)
		_, err := r.client.Get(ctx, key).Result()
		if err == redis.Nil {
			if err := r.client.Set(ctx, key, 0, 0).Err(); err != nil {
				return err
			}
		}
	}
	return nil
}
