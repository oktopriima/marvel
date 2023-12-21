package mysqlrepo

import (
	"context"
	"errors"
	"github.com/fatih/structs"
	"github.com/oktopriima/marvel/app/modules/base/filter"
	"github.com/oktopriima/marvel/app/modules/base/model"
	"github.com/oktopriima/marvel/app/modules/base/repo"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{
		db: db,
	}
}

func (r *BaseRepo) GetDB(ctx context.Context) *gorm.DB {
	db := r.db
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db = db.WithContext(newCtx)

	return db
}

func (r *BaseRepo) FindByID(ctx context.Context, m model.Model, id int64, preloadFields ...string) error {
	q := r.GetDB(ctx)

	for _, p := range preloadFields {
		q = q.Preload(p)
	}

	err := q.Where("id = ?", id).Take(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repo.RecordNotFound
	}
	return err
}

func (r *BaseRepo) CreateOrUpdate(ctx context.Context, m model.Model, query interface{}, attrs ...interface{}) error {
	return r.GetDB(ctx).Where(query).Assign(attrs...).FirstOrCreate(m).Error
}

func (r *BaseRepo) Update(ctx context.Context, m model.Model, attrs ...interface{}) error {
	return r.GetDB(ctx).Model(m).Updates(toSearchableMap(attrs...)).Error
}

func (r *BaseRepo) Updates(ctx context.Context, m model.Model, params interface{}) error {
	return r.GetDB(ctx).Model(m).Updates(params).Error
}

func (r *BaseRepo) Create(ctx context.Context, m model.Model) error {
	return r.GetDB(ctx).Create(m).Error
}

func (r *BaseRepo) Search(ctx context.Context, val interface{}, f filter.Filter, preloadFields ...string) error {
	q := r.GetDB(ctx).Model(val)
	for query, val := range f.GetWhere() {
		q = q.Where(query, val...)
	}

	for query, val := range f.GetJoins() {
		q = q.Joins(query, val...)
	}

	if f.GetGroups() != "" {
		q = q.Group(f.GetGroups())
	}

	if f.GetLimit() > 0 {
		q = q.Limit(f.GetLimit())
	}

	if len(f.GetOrderBy()) > 0 {
		for _, order := range f.GetOrderBy() {
			q = q.Order(order)
		}
	}

	for _, p := range preloadFields {
		q = q.Preload(p)
	}

	return q.Offset(f.GetOffset()).Find(val).Error
}

func (r *BaseRepo) Save(ctx context.Context, m model.Model) error {
	return r.GetDB(ctx).Model(m).Save(m).Error
}

func (r *BaseRepo) DeleteByID(ctx context.Context, m model.Model, id int64) error {
	db := r.GetDB(ctx).Where("id = ?", id).Take(m)
	if db.Error != nil || m.GetID() == 0 {
		return repo.RecordNotFound
	}
	return db.Delete(m).Error
}

func toSearchableMap(attrs ...interface{}) (result interface{}) {
	if len(attrs) > 1 {
		if str, ok := attrs[0].(string); ok {
			result = map[string]interface{}{str: attrs[1]}
		}
	} else if len(attrs) == 1 {
		if attr, ok := attrs[0].(map[string]interface{}); ok {
			result = attr
		}

		if attr, ok := attrs[0].(interface{}); ok {
			s := structs.New(attr)
			s.TagName = "json"
			m := s.Map()

			value := make(map[string]interface{}, len(m))
			var ns schema.NamingStrategy
			for col, val := range m {
				dbCol := ns.ColumnName("", col)
				value[dbCol] = val
			}
			result = value
		}
	}
	return
}
