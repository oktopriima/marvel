package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/src/app/entity/models"
	. "gopkg.in/check.v1"
	"testing"
)

type S struct {
	instance database.DBInstance
	mock     sqlmock.Sqlmock
}

func Test(t *testing.T) {
	TestingT(t)
}

var dbInstance, mock = Instance()

var _ = Suite(&S{
	instance: dbInstance,
	mock:     mock,
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
