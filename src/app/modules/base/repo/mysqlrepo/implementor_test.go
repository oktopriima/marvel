package mysqlrepo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
	"github.com/oktopriima/marvel/src/test"
	. "gopkg.in/check.v1"
	"gorm.io/gorm"
	"testing"
	"time"
)

var databaseInstance, sqlMock = test.Instance()

type S struct {
	instance database.DBInstance
	mock     sqlmock.Sqlmock
	ctx      context.Context
}

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{
	instance: databaseInstance,
	mock:     sqlMock,
	ctx:      context.Background(),
})

type MysqlTestExample struct {
	model.BaseModel
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	DeletedAt gorm.DeletedAt `gorm:"default:null" json:"deleted_at"`
}

var examples []*MysqlTestExample

func (s *S) Factory(examples []*MysqlTestExample) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"id", "name", "age", "created_at", "updated_at", "deleted_at",
	})

	for _, data := range examples {
		rows.AddRow(data.Id, data.Name, data.Age, data.CreatedAt, data.UpdatedAt, nil)
	}

	return rows
}

func (s *S) TestFindById(c *C) {
	examples = append(examples, &MysqlTestExample{
		BaseModel: model.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "john doe",
		Age:  20,
	})

	rows := s.Factory(examples)

	expectedQuery := "SELECT * FROM `mysql_test_examples` WHERE id = ? AND `mysql_test_examples`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnRows(rows)

	repo := NewBaseMysqlRepo(s.instance)

	var result MysqlTestExample
	err := repo.FindByID(s.ctx, &result, 1)
	c.Assert(err, IsNil)
	c.Assert(result.Id, Equals, int64(1))
	c.Assert(result.Name, Equals, "john doe")
	c.Assert(result.Age, Equals, 20)
}

func (s *S) TestCreate(c *C) {
	now := time.Now()
	expectedQuery := "INSERT INTO `mysql_test_examples` (`created_at`,`updated_at`,`name`,`age`) VALUES (?,?,?,?)"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(expectedQuery).
		WithArgs(now, now, "john doe dune", 20).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	repo := NewBaseMysqlRepo(s.instance)
	err := repo.Create(s.ctx, &MysqlTestExample{
		BaseModel: model.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name: "john doe dune",
		Age:  20,
	})
	c.Assert(err, IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) TestUpdate(c *C) {
	now := time.Now().Round(time.Millisecond)

	m := MysqlTestExample{
		BaseModel: model.BaseModel{
			Id:        1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name: "john doe",
		Age:  20,
	}

	expectedQuery := "UPDATE `mysql_test_examples` SET `created_at`=?,`updated_at`=?,`name`=?,`age`=? WHERE `mysql_test_examples`.`deleted_at` IS NULL AND `id` = ?"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(expectedQuery).
		WithArgs(now, now, "john doe dune", 20, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	repo := NewBaseMysqlRepo(s.instance)

	// update the model
	m.Name = "john doe dune"
	err := repo.Update(s.ctx, &m)

	c.Assert(err, IsNil)
	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) TestSave(c *C) {
	now := time.Now()
	expectedQuery := "INSERT INTO `mysql_test_examples` (`created_at`,`updated_at`,`name`,`age`) VALUES (?,?,?,?)"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(expectedQuery).
		WithArgs(now, now, "john doe dune", 20).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	repo := NewBaseMysqlRepo(s.instance)
	err := repo.Save(s.ctx, &MysqlTestExample{
		BaseModel: model.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name: "john doe dune",
		Age:  20,
	})
	c.Assert(err, IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) TestDelete(c *C) {
	now := time.Now().Round(time.Second)

	m := MysqlTestExample{
		BaseModel: model.BaseModel{
			Id:        1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name: "john doe",
		Age:  20,
	}

	examples = append(examples, &m)
	rows := s.Factory(examples)

	expectedSelectQuery := "SELECT * FROM `mysql_test_examples` WHERE id = ? AND `mysql_test_examples`.`deleted_at` IS NULL AND `mysql_test_examples`.`id` = ? LIMIT ?"
	s.mock.ExpectQuery(expectedSelectQuery).
		WithArgs(1, 1, 1).
		WillReturnRows(rows)

	expectedDeleteQuery := "UPDATE `mysql_test_examples` SET `deleted_at`=? WHERE id = ? AND `mysql_test_examples`.`deleted_at` IS NULL AND `mysql_test_examples`.`id` = ? AND `mysql_test_examples`.`id` = ? LIMIT ?"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(expectedDeleteQuery).
		WithArgs(sqlmock.AnyArg(), 1, 1, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	repo := NewBaseMysqlRepo(s.instance)
	err := repo.DeleteByID(s.ctx, &m, m.Id)
	c.Assert(err, IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}
