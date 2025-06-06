package contract

import (
	"context"
	"errors"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
)

var RecordNotFound = errors.New("record not found")

type BaseMysqlRepo interface {
	Updatable
	Saveable
	Creatable
	Removable
	CanFindByID
}

type Updatable interface {
	Update(ctx context.Context, m models.Model, attrs ...interface{}) error
}

type Saveable interface {
	Save(ctx context.Context, m models.Model) error
}

type Creatable interface {
	Create(ctx context.Context, m models.Model) error
}

type Removable interface {
	DeleteByID(ctx context.Context, m models.Model, id int64) error
}

type CanFindByID interface {
	FindByID(ctx context.Context, m models.Model, id int64, preloadFields ...string) error
}
