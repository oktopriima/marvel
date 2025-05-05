package mysqlrepo

import (
	"context"
	"errors"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/pkg/util"
	"github.com/oktopriima/marvel/src/app/modules/base/contract"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
	"go.elastic.co/apm/v2"
	"gorm.io/gorm"
)

type BaseMysqlRepo struct {
	db *gorm.DB
}

func NewBaseMysqlRepo(mysql database.DBInstance) *BaseMysqlRepo {
	return &BaseMysqlRepo{
		db: mysql.Database(),
	}
}

func (r *BaseMysqlRepo) GetDB(ctx context.Context) *gorm.DB {
	db := r.db

	db = db.WithContext(ctx)

	return db
}

func (r *BaseMysqlRepo) FindByID(ctx context.Context, m models.Model, id int64, preloadFields ...string) error {
	span, ctx := apm.StartSpan(ctx, "mysqlRepo.FindByID", tracer.RepositoryTraceName)
	defer span.End()

	q := r.GetDB(ctx)

	for _, p := range preloadFields {
		q = q.Preload(p)
	}

	err := q.Where("id = ?", id).Take(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return contract.RecordNotFound
	}
	return err
}

func (r *BaseMysqlRepo) Update(ctx context.Context, m models.Model, attrs ...interface{}) error {
	span, ctx := apm.StartSpan(ctx, "mysqlRepo.Update", tracer.RepositoryTraceName)
	defer span.End()

	return r.GetDB(ctx).Model(m).Updates(util.ToSearchableMap(attrs...)).Error
}

func (r *BaseMysqlRepo) Create(ctx context.Context, m models.Model) error {
	span, ctx := apm.StartSpan(ctx, "mysqlRepo.Create", tracer.RepositoryTraceName)
	defer span.End()

	return r.GetDB(ctx).Create(m).Error
}

func (r *BaseMysqlRepo) Save(ctx context.Context, m models.Model) error {
	span, ctx := apm.StartSpan(ctx, "mysqlRepo.Save", tracer.RepositoryTraceName)
	defer span.End()

	return r.GetDB(ctx).Model(m).Save(m).Error
}

func (r *BaseMysqlRepo) DeleteByID(ctx context.Context, m models.Model, id int64) error {
	span, ctx := apm.StartSpan(ctx, "mysqlRepo.DeleteByID", tracer.RepositoryTraceName)
	defer span.End()

	db := r.GetDB(ctx).Where("id = ?", id).Take(m)
	if db.Error != nil || m.GetID() == 0 {
		return contract.RecordNotFound
	}
	return db.Delete(m).Error
}
