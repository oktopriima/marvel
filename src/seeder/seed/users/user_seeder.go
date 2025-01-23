package users

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/helper"
	"github.com/oktopriima/marvel/src/app/modules/base/model"
	"github.com/oktopriima/marvel/src/seeder/implementer"
	"gorm.io/gorm"
)

var userData = []models.Users{
	{
		Email:     "octoprima93@gmail.com",
		FirstName: "okto",
		LastName:  "prima",
		Password:  helper.GeneratePassword("delicious"),
	},
}

func Run(ctx context.Context, db *gorm.DB) {
	var m models.Users
	s := implementer.NewSeederImplementer(db)
	if check := s.CheckRow(ctx, &m); !check {
		return
	}

	var data []model.Model

	for _, datum := range userData {
		data = append(data, &datum)
	}

	err := s.Run(ctx, data)
	if err != nil {
		return
	}
}
