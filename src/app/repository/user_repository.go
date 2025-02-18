package repository

import (
	"context"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/modules/base/repo/mysqlrepo"
	"github.com/oktopriima/marvel/src/app/modules/base/repo/redisrepo"
	"github.com/oktopriima/marvel/src/app/repository/contract"
	"go.elastic.co/apm/v2"
)

type userRepository struct {
	*mysqlrepo.BaseMysqlRepo
	*redisrepo.BaseRedisRepo
}

func NewUserRepository(
	mysqlInstance database.DBInstance,
	redisInstance cache.RedisInstance,
) contract.UserContract {
	return &userRepository{
		BaseMysqlRepo: mysqlrepo.NewBaseMysqlRepo(mysqlInstance),
		BaseRedisRepo: redisrepo.NewBaseRedisRepo(redisInstance),
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

func (u *userRepository) GetByEmail(email string, ctx context.Context) ([]*models.Users, error) {
	span, ctx := apm.StartSpan(ctx, "userRepository.GetByEmail", tracer.RepositoryTraceName)
	defer span.End()

	var users []*models.Users
	db := u.GetDB(ctx)

	if err := db.Where("email LIKE '%?%'", email).Find(users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
