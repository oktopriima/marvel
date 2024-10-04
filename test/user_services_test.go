package test_test

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/helper"
	"github.com/oktopriima/marvel/app/modules/base/model"
	"github.com/oktopriima/marvel/app/repository"
	. "gopkg.in/check.v1"
	"gorm.io/gorm"
	"time"
)

func (s *S) Test_userServices_Successful_FindByEmail(c *C) {
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

	expectedQuery := "SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs("jhon@gmail.com", 1).
		WillReturnRows(rows)

	userServices := repository.NewUserRepository(s.instance)
	user, err := userServices.FindByEmail("jhon@gmail.com", context.Background())

	c.Assert(err, IsNil)

	c.Assert(user, NotNil)
	c.Assert(user.Id, Equals, int64(1))
	c.Assert(user.FirstName, Equals, "jhon")
	c.Assert(user.LastName, Equals, "doe")
	c.Assert(user.Email, Equals, "jhon@gmail.com")

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) Test_userServices_Failed_FindByEmail(c *C) {
	expectedQuery := "SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs("doe@gmail.com", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	userServices := repository.NewUserRepository(s.instance)
	user, err := userServices.FindByEmail("doe@gmail.com", context.Background())

	c.Assert(err, NotNil)
	c.Assert(user, IsNil)
	c.Assert(err.Error(), Equals, "record not found")

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}
