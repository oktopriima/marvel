package models

import (
	"github.com/oktopriima/marvel/app/modules/base/model"
	"time"
)

type Users struct {
	model.BaseModel
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"password"`
	RememberToken   string    `json:"remember_token"`
	DeletedAt       time.Time `gorm:"default:null" json:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
