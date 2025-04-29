package user_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/helper"
	baseModel "github.com/oktopriima/marvel/src/app/modules/base/models"
	"github.com/oktopriima/marvel/src/app/repository"
	uc "github.com/oktopriima/marvel/src/app/usecase/users"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
	. "gopkg.in/check.v1"
	"gorm.io/gorm"
	"time"
)

func (s *S) Test_User_Usecase_Successful_FindById(c *C) {
	m := models.Users{
		BaseModel: baseModel.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:      "jhon",
		Email:     "jhon@gmail.com",
		Password:  helper.GeneratePassword("password123"),
		DeletedAt: gorm.DeletedAt{},
	}
	users = append(users, &m)
	rows := s.UserFactory(users)

	key := fmt.Sprintf("%s:%d", m.TableName(), m.Id)
	s.redisMock.ExpectGet(key).RedisNil()

	expectedQuery := "SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnRows(rows)

	bytes, err := json.Marshal(m)
	c.Assert(err, IsNil)
	s.redisMock.ExpectSet(key, string(bytes), 10*time.Hour).SetVal("OK")

	expectation := new(dto.UserResponse)
	expectation.Id = users[0].Id
	expectation.FullName = users[0].Name

	repo := repository.NewUserRepository(s.instance, s.redisInstance)
	userUsecase := uc.NewUserUsecase(repo)

	ctx := context.Background()

	resp, err := userUsecase.FindByID(ctx, int64(1))
	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)
	c.Assert(resp, DeepEquals, expectation)

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) Test_User_Usecase_Successful_FindById_OnCache(c *C) {
	m := models.Users{
		BaseModel: baseModel.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:      "jhon",
		Email:     "jhon@gmail.com",
		Password:  helper.GeneratePassword("password123"),
		DeletedAt: gorm.DeletedAt{},
	}

	byteModel, err := json.Marshal(m)
	c.Assert(err, IsNil)

	key := fmt.Sprintf("%s:%d", m.TableName(), m.Id)
	s.redisMock.ExpectGet(key).SetVal(string(byteModel))

	repo := repository.NewUserRepository(s.instance, s.redisInstance)
	userUsecase := uc.NewUserUsecase(repo)
	resp, err := userUsecase.FindByID(s.ctx, m.Id)
	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)
	c.Assert(resp.Id, Equals, m.Id)
	c.Assert(resp.FullName, Equals, m.Name)
}

func (s *S) Test_User_Usecase_Failed_FindById(c *C) {
	m := new(models.Users)
	m.Id = 1

	key := fmt.Sprintf("%s:%d", m.TableName(), m.Id)
	s.redisMock.ExpectGet(key).RedisNil()

	expectedQuery := "SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	repo := repository.NewUserRepository(s.instance, s.redisInstance)
	userUsecase := uc.NewUserUsecase(repo)

	ctx := context.Background()
	resp, err := userUsecase.FindByID(ctx, int64(1))

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "record not found")
	c.Assert(resp, IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}
