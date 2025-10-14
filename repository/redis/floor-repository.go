package redis




import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)





type FloorRepo struct {
	client *redis.Client
	ctx    context.Context
}


func NewFloorRepo(client *redis.Client) *FloorRepo {
	return &FloorRepo{
		client: client,
		ctx:    context.Background(),
	}
}


func (r *FloorRepo) IncreaseFloorCount(floor int) error {
	key := fmt.Sprintf("floor:%d", floor)
	return r.client.Incr(r.ctx, key).Err()
}

func (r *FloorRepo) DecreaseFloorCount(floor int) error {
	key := fmt.Sprintf("floor:%d", floor)
	return r.client.Decr(r.ctx, key).Err()
}

func (r *FloorRepo) GetFloorCount(floor int) (int, error) {
	key := fmt.Sprintf("floor:%d", floor)
	val, err := r.client.Get(r.ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}


func (r *FloorRepo) InitializeFloorCounts(totalFloors int) error {
	for i := 1; i <= totalFloors; i++ {
		key := fmt.Sprintf("floor:%d", i)
		_, err := r.client.Get(r.ctx, key).Result()
		if err == redis.Nil {
			if err := r.client.Set(r.ctx, key, 0, 0).Err(); err != nil {
				return err
			}
		}
	}
	return nil
}