package user_test

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/helper"
	baseModel "github.com/oktopriima/marvel/src/app/modules/base/models"
	"github.com/oktopriima/marvel/src/app/repository"
	. "gopkg.in/check.v1"
	"gorm.io/gorm"
	"time"
)

func (s *S) Test_userServices_Successful_FindByEmail(c *C) {
	users = append(users, &models.Users{
		BaseModel: baseModel.BaseModel{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:     "jhon@gmail.com",
		Password:  helper.GeneratePassword("password123"),
		DeletedAt: gorm.DeletedAt{},
	})
	rows := s.UserFactory(users)

	expectedQuery := "SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs("jhon@gmail.com", 1).
		WillReturnRows(rows)

	userServices := repository.NewUserRepository(s.instance, s.redisInstance)
	user, err := userServices.FindByEmail("jhon@gmail.com", context.Background())

	c.Assert(err, IsNil)

	c.Assert(user, NotNil)
	c.Assert(user.Id, Equals, int64(1))
	c.Assert(user.Email, Equals, "jhon@gmail.com")

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}

func (s *S) Test_userServices_Failed_FindByEmail(c *C) {
	expectedQuery := "SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT ?"
	s.mock.ExpectQuery(expectedQuery).
		WithArgs("doe@gmail.com", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	userServices := repository.NewUserRepository(s.instance, s.redisInstance)
	user, err := userServices.FindByEmail("doe@gmail.com", context.Background())

	c.Assert(err, NotNil)
	c.Assert(user, IsNil)
	c.Assert(err.Error(), Equals, "record not found")

	c.Assert(s.mock.ExpectationsWereMet(), IsNil)
}
