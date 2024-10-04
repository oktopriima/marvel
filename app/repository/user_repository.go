package repository

import (
	"context"
	"github.com/oktopriima/marvel/app/contract"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/modules/base/repo/mysqlrepo"
	"github.com/oktopriima/marvel/core/database"
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
	user := new(models.Users)
	db := u.GetDB(ctx)

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
