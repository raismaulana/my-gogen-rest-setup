package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/database"
)

type RedisGateway struct {
	database.RedisClientImpl
}

func NewRedisGateway(rdb *redis.Client) *RedisGateway {
	return &RedisGateway{
		RedisClientImpl: database.RedisClientImpl{
			RDB: rdb,
		},
	}
}

func (r *RedisGateway) RDBSet(ctx context.Context, RDBkey string, value interface{}, expiration time.Duration) error {
	return r.RDB.Set(ctx, RDBkey, value, expiration).Err()
}

func (r *RedisGateway) RDBGet(ctx context.Context, RDBkey string) (string, error) {
	return r.RDB.Get(ctx, RDBkey).Result()
}
