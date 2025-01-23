package seed

import (
	"context"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
	"gorm.io/gorm"
)

type BaseSeeder struct {
	db *gorm.DB
}

func NewBaseSeeder(db *gorm.DB) *BaseSeeder {
	return &BaseSeeder{
		db: db,
	}
}

func (b *BaseSeeder) GetDB(ctx context.Context) *gorm.DB {
	db := b.db

	db = db.WithContext(ctx)

	return db
}

func (b *BaseSeeder) CheckRow(ctx context.Context, m model.Model) bool {
	var count int64
	b.GetDB(ctx).Model(m).Count(&count)

	if count > 0 {
		return false
	}
	return true
}

func (b *BaseSeeder) Run(ctx context.Context, m []model.Model) error {
	tx := b.GetDB(ctx).Begin()

	defer func() {
		tx.Rollback()
	}()

	for _, data := range m {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
