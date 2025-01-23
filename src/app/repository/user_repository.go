package repository

import (
	"context"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/modules/base/repo/mysqlrepo"
	"github.com/oktopriima/marvel/src/app/repository/contract"
	"go.elastic.co/apm/v2"
)

type userRepository struct {
	*mysqlrepo.BaseRepo
}

func NewUserRepository(instance database.DBInstance) contract.UserContract {
	return &userRepository{
		BaseRepo: mysqlrepo.NewBaseRepo(instance),
	}
}

func (u *userRepository) FindByEmail(email string, ctx context.Context) (*models.Users, error) {
	span, ctx := apm.StartSpan(ctx, "userRepository.FindByEmail", tracer.RepositoryTraceName)
	defer span.End()

	user := new(models.Users)
	db := u.GetDB(ctx)

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
