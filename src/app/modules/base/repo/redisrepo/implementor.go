package redisrepo

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
	"time"
)

type BaseRedisRepo struct {
	redis *redis.Pool
}

func NewBaseRedisRepo(instance cache.RedisInstance) *BaseRedisRepo {
	return &BaseRedisRepo{
		redis: instance.Database(),
	}
}

func (r *BaseRedisRepo) FindCache(ctx context.Context, key string) (model.Model, error) {
	return nil, nil
}

func (r *BaseRedisRepo) FindRawCache(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (r *BaseRedisRepo) StoreCache(ctx context.Context, key string, ttl time.Duration, m model.Model) error {
	return nil
}

func (r *BaseRedisRepo) StoreObjectCache(ctx context.Context, key string, ttl time.Duration, m []byte) error {
	return nil
}

func (r *BaseRedisRepo) RemoveCache(ctx context.Context, key string) error {
	return nil
}
