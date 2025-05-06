package users

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/helper"
	baseModel "github.com/oktopriima/marvel/src/app/modules/base/models"
	"github.com/oktopriima/marvel/src/cmd/seeder/implementer"
	"gorm.io/gorm"
	"time"
)

var userData = []models.Users{
	{
		Name:            "john doe",
		Email:           "octoprima93@gmail.com",
		Password:        helper.GeneratePassword("delicious"),
		EmailVerifiedAt: time.Now(),
	},
}

func Run(ctx context.Context, db *gorm.DB) {
	var m models.Users
	s := implementer.NewSeederImplementer(db)
	if check := s.CheckRow(ctx, &m); !check {
		return
	}

	var data []baseModel.Model

	for _, datum := range userData {
		data = append(data, &datum)
	}

	err := s.Run(ctx, data)
	if err != nil {
		return
	}
}
