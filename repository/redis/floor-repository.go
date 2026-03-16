package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
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
	key := fmt.Sprintf("floor:%d", floor)
	return r.client.Incr(ctx, key).Err()
}

func (r *FloorRepo) DecreaseFloorCount(ctx context.Context, floor int) error {
	key := fmt.Sprintf("floor:%d", floor)
	return r.client.Decr(ctx, key).Err()
}

func (r *FloorRepo) GetFloorCount(ctx context.Context, floor int) (int, error) {
	key := fmt.Sprintf("floor:%d", floor)
	val, err := r.client.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}

func (r *FloorRepo) InitializeFloorCounts(ctx context.Context, totalFloors int) error {

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
