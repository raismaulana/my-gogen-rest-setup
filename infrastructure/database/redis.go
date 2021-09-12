package database

import "github.com/go-redis/redis/v8"

type RedisClientImpl struct {
	RDB *redis.Client
}
