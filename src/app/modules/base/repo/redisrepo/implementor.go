package redisrepo

import (
	"context"
	"encoding/json"
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

func (r *BaseRedisRepo) FindCache(ctx context.Context, m model.Model, key string) (model.Model, error) {
	conn, err := r.redis.GetContext(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	bytes, err := redis.Bytes(conn.Do("GET", key))
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
	conn, err := r.redis.GetContext(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	bytes, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (r *BaseRedisRepo) StoreCache(ctx context.Context, key string, ttl time.Duration, m model.Model) error {
	conn, err := r.redis.GetContext(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Do("SETEX", key, ttl, m)
	return err
}

func (r *BaseRedisRepo) StoreObjectCache(ctx context.Context, key string, ttl time.Duration, m []byte) error {
	conn, err := r.redis.GetContext(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Do("SETEX", key, ttl, m)
	return err
}

func (r *BaseRedisRepo) RemoveCache(ctx context.Context, key string) error {
	conn, err := r.redis.GetContext(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()
	_, err = conn.Do("DEL", key)
	return err
}
