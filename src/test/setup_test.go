package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/src/app/entity/models"
	. "gopkg.in/check.v1"
	"testing"
)

type S struct {
	instance  database.DBInstance
	redis     cache.RedisInstance
	mock      sqlmock.Sqlmock
	redisMock redismock.ClientMock
}

func Test(t *testing.T) {
	TestingT(t)
}

var dbInstance, mock = Instance()
var redisInstance, redisMock = RedisInstance()

var _ = Suite(&S{
	instance:  dbInstance,
	redis:     redisInstance,
	mock:      mock,
	redisMock: redisMock,
})

var users []*models.Users

func (s *S) InsertUserData(users []*models.Users) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"id", "email", "password", "created_at", "updated_at", "deleted_at",
	})

	for _, user := range users {
		rows.AddRow(user.Id, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, nil)
	}

	return rows
}
