package redisrepo

import (
	"context"
	"encoding/json"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
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

func (r *BaseRedisRepo) FindCache(ctx context.Context, m model.Model, key string) (model.Model, error) {
	conn := r.redis.Conn()
	defer conn.Close()

	bytes, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *BaseRedisRepo) FindRawCache(ctx context.Context, key string) ([]byte, error) {
	conn := r.redis.Conn()
	defer conn.Close()

	bytes, err := conn.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (r *BaseRedisRepo) StoreCache(ctx context.Context, key string, ttl time.Duration, m model.Model) error {
	conn := r.redis.Conn()
	defer conn.Close()

	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := conn.Set(ctx, key, string(marshal), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r *BaseRedisRepo) StoreObjectCache(ctx context.Context, key string, ttl time.Duration, m []byte) error {
	conn := r.redis.Conn()
	defer conn.Close()

	err := conn.Set(ctx, key, string(m), ttl).Err()
	return err
}

func (r *BaseRedisRepo) RemoveCache(ctx context.Context, key string) error {
	conn := r.redis.Conn()
	defer conn.Close()
	return conn.Del(ctx, key).Err()
}
