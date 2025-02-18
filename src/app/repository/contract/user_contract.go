package contract

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/modules/base/contract"
)

type UserContract interface {
	contract.BaseMysqlRepo
	contract.BaseRedisRepo
	SearchByEmail
}

type SearchByEmail interface {
	FindByEmail(email string, ctx context.Context) (*models.Users, error)
	GetByEmail(email string, ctx context.Context) ([]*models.Users, error)
}
