package models

import (
	"github.com/oktopriima/marvel/app/modules/base/model"
	"gorm.io/gorm"
)

type Users struct {
	model.BaseModel
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Password     string         `json:"password"`
	RefreshToken string         `json:"refresh_token"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
