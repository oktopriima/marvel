package setup

import (
	"github.com/DATA-DOG/go-sqlmock"
	core "github.com/oktopriima/marvel/core/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed to open sql mock: " + err.Error())
	}

	mock.ExpectQuery(`SELECT VERSION()`).WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm DB:" + err.Error())
	}

	return gormDB, mock
}

func SetupInstance() (core.DBInstance, sqlmock.Sqlmock) {
	gormDb, mock := setupMock()
	ins := new(core.Instance)
	ins.GormDB = gormDb

	return ins, mock
}
