package contract

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/modules/base/repo"
)

type UserContract interface {
	repo.BaseRepo
	EmailLogin
}

type EmailLogin interface {
	FindByEmail(email string, ctx context.Context) (*models.Users, error)
}
