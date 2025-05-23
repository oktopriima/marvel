package contract

import (
	"context"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
	"time"
)

type BaseRedisRepo interface {
	Cacheable
	CacheRemovable
	CanFindCache
}

type Cacheable interface {
	StoreCache(ctx context.Context, key string, ttl time.Duration, m models.Model) error
	StoreObjectCache(ctx context.Context, key string, ttl time.Duration, m []byte) error
}

type CacheRemovable interface {
	RemoveCache(ctx context.Context, key string) error
}

type CanFindCache interface {
	FindCache(ctx context.Context, m models.Model, key string) error
	FindRawCache(ctx context.Context, key string) ([]byte, error)
}
