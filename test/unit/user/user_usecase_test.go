package user_test

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/helper"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
	"github.com/oktopriima/marvel/src/app/repository"
	userUsecase "github.com/oktopriima/marvel/src/app/usecase/users"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
	"gopkg.in/check.v1"
	"gorm.io/gorm"
	"time"
)

func (s *S) Test_User_Usecase_Successful_FindById(c *check.C) {
	users = append(users, &models.Users{
		BaseModel: model.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:     "jhon@gmail.com",
		Name:      "jhon",
		Password:  helper.GeneratePassword("password123"),
		DeletedAt: gorm.DeletedAt{},
	})
	rows := s.UserFactory(users)

	expectedQuery := "SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnRows(rows)

	expectation := new(dto.UserResponse)
	expectation.Id = users[0].Id
	expectation.FullName = fmt.Sprintf("%s", users[0].Name)

	repo := repository.NewUserRepository(s.instance)
	usecase := userUsecase.NewUserUsecase(repo)

	ctx := context.Background()

	resp, err := usecase.FindByID(ctx, int64(1))
	c.Assert(err, check.IsNil)
	c.Assert(resp, check.NotNil)
	c.Assert(resp, check.DeepEquals, expectation)

	c.Assert(s.mock.ExpectationsWereMet(), check.IsNil)
}

func (s *S) Test_User_Usecase_Failed_FindById(c *check.C) {
	expectedQuery := "SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	repo := repository.NewUserRepository(s.instance)
	usecase := userUsecase.NewUserUsecase(repo)

	ctx := context.Background()
	resp, err := usecase.FindByID(ctx, int64(1))

	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "record not found")
	c.Assert(resp, check.IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), check.IsNil)
}
