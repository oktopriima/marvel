package test_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/core/database"
	"github.com/oktopriima/marvel/test/setup"
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

var dbInstance, mock = setup.SetupInstance()

var _ = Suite(&S{
	instance: dbInstance,
	mock:     mock,
})

var usersData []*models.Users

func (s *S) InsertUserData(users []*models.Users) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"id", "email", "first_name", "last_name", "password", "refresh_token", "created_at", "updated_at", "deleted_at",
	})

	for _, user := range users {
		rows.AddRow(user.Id, user.Email, user.FirstName, user.LastName, user.Password, user.RefreshToken, user.CreatedAt, user.UpdatedAt, nil)
	}

	return rows
}
