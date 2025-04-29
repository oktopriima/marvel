package dto

import (
	"github.com/oktopriima/marvel/src/app/entity/models"
)

type UserResponse struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (r UserResponse) ConvertToResponse(users *models.Users) *UserResponse {
	return &UserResponse{
		Id:       users.Id,
		FullName: users.Name,
	}
}
