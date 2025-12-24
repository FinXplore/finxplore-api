package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository interface {
	SetSession(key string, value string, ttl time.Duration) error
	GetSession(key string) (string, error)
	DeleteSession(key string) error
}

type redisCache struct {
	rdb *redis.Client
}

func NewCacheRepository(rdb *redis.Client) CacheRepository {
	return &redisCache{rdb: rdb}
}

// --- Implementation ---

func (r *redisCache) SetSession(key string, value string, ttl time.Duration) error {
	ctx := context.Background()
	return r.rdb.Set(ctx, key, value, ttl).Err()
}

func (r *redisCache) GetSession(key string) (string, error) {
	ctx := context.Background()
	return r.rdb.Get(ctx, key).Result()
}

func (r *redisCache) DeleteSession(key string) error {
	ctx := context.Background()
	return r.rdb.Del(ctx, key).Err()
}