package models

import (
	"github.com/oktopriima/marvel/src/app/modules/base/models"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	models.BaseModel
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	Password        string         `json:"password"`
	RememberToken   string         `json:"remember_token"`
	DeletedAt       gorm.DeletedAt `gorm:"default:null" json:"deleted_at"`
}

func (u *Users) TableName() string {
	return "users"
}
