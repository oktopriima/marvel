package user_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/test"
	. "gopkg.in/check.v1"
	"testing"
)

type S struct {
	instance      database.DBInstance
	redisInstance cache.RedisInstance
	mock          sqlmock.Sqlmock
	redisMock     redismock.ClientMock
	ctx           context.Context
}

func Test(t *testing.T) {
	TestingT(t)
}

var dbInstance, mock = test.Instance()
var redisInstance, redisMock = test.RedisInstance()

var _ = Suite(&S{
	instance:      dbInstance,
	redisInstance: redisInstance,
	mock:          mock,
	redisMock:     redisMock,
	ctx:           context.Background(),
})

var users []*models.Users

func (s *S) UserFactory(users []*models.Users) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"id", "email", "name", "password", "created_at", "updated_at", "deleted_at",
	})

	for _, user := range users {
		rows.AddRow(user.Id, user.Email, user.Name, user.Password, user.CreatedAt, user.UpdatedAt, nil)
	}

	return rows
}
