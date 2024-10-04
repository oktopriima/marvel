package test_test

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/helper"
	"github.com/oktopriima/marvel/app/modules/base/model"
	"github.com/oktopriima/marvel/app/repository"
	"github.com/oktopriima/marvel/app/usecase/users"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
	"gopkg.in/check.v1"
	"gorm.io/gorm"
	"time"
)

func (s *S) Test_User_Usecase_Successful_FindById(c *check.C) {
	usersData = append(usersData, &models.Users{
		BaseModel: model.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:        "jhon@gmail.com",
		FirstName:    "jhon",
		LastName:     "doe",
		Password:     helper.GeneratePassword("password123"),
		RefreshToken: "",
		DeletedAt:    gorm.DeletedAt{},
	})
	rows := s.InsertUserData(usersData)

	expectedQuery := "SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs(1, 1).
		WillReturnRows(rows)

	expectation := new(dto.UserResponse)
	expectation.Id = usersData[0].Id
	expectation.FullName = fmt.Sprintf("%s %s", usersData[0].FirstName, usersData[0].LastName)

	repo := repository.NewUserRepository(s.instance)
	userUsecase := users.NewUserUsecase(repo)

	ctx := context.Background()

	resp, err := userUsecase.FindByID(ctx, int64(1))
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
	userUsecase := users.NewUserUsecase(repo)

	ctx := context.Background()
	resp, err := userUsecase.FindByID(ctx, int64(1))

	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "record not found")
	c.Assert(resp, check.IsNil)

	c.Assert(s.mock.ExpectationsWereMet(), check.IsNil)
}
