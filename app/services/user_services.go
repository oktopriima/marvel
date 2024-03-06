package services

import (
	"github.com/oktopriima/marvel/app/modules/base/repo/mysqlrepo"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/core/database"
)

type userServices struct {
	*mysqlrepo.BaseRepo
}

func NewUserServices(instance database.DBInstance) repository.UserRepository {
	return &userServices{
		BaseRepo: mysqlrepo.NewBaseRepo(instance),
	}
}
