package services

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/modules/base/repo/mysqlrepo"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/core/database"
)

type userServices struct {
	*mysqlrepo.BaseRepo
}

func (u *userServices) FindByEmail(email string, ctx context.Context) (*models.Users, error) {
	user := new(models.Users)
	db := u.GetDB(ctx)

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserServices(instance database.DBInstance) repository.UserRepository {
	return &userServices{
		BaseRepo: mysqlrepo.NewBaseRepo(instance),
	}
}
