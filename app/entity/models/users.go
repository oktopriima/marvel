package models

import (
	"github.com/oktopriima/marvel/app/modules/base/model"
	"gorm.io/gorm"
)

type Users struct {
	model.BaseModel
	FirstName  string         `json:"first_name"`
	MiddleName string         `json:"middle_name"`
	LastName   string         `json:"last_name"`
	Nickname   string         `json:"nickname"`
	Password   string         `json:"password"`
	Avatar     string         `json:"avatar"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
