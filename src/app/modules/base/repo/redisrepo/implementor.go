package redisrepo

import (
	"context"
	"encoding/json"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type BaseRedisRepo struct {
	redis *redis.Client
}

func NewBaseRedisRepo(instance cache.RedisInstance) *BaseRedisRepo {
	return &BaseRedisRepo{
		redis: instance.Database(),
	}
}

func (r *BaseRedisRepo) FindCache(ctx context.Context, m models.Model, key string) error {
	str, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(str, &m)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRedisRepo) FindRawCache(ctx context.Context, key string) ([]byte, error) {
	bytes, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (r *BaseRedisRepo) StoreCache(ctx context.Context, key string, ttl time.Duration, m models.Model) error {
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := r.redis.Set(ctx, key, string(marshal), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r *BaseRedisRepo) StoreObjectCache(ctx context.Context, key string, ttl time.Duration, m []byte) error {
	err := r.redis.Set(ctx, key, string(m), ttl).Err()
	return err
}

func (r *BaseRedisRepo) RemoveCache(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}
