package dto

import (
	"github.com/oktopriima/marvel/src/app/entity/models"
	"time"
)

type UserResponse interface {
	GetObject() *userResponse
}

type userResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *userResponse) GetObject() *userResponse {
	return u
}

func ConvertToResponse(users *models.Users) UserResponse {
	return &userResponse{
		Id:    users.Id,
		Name:  users.Name,
		Email: users.Email,
	}
}

type NotifyLoginRequest struct {
	Id              int64       `json:"id"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Name            string      `json:"name"`
	Email           string      `json:"email"`
	EmailVerifiedAt time.Time   `json:"email_verified_at"`
	Password        string      `json:"password"`
	RememberToken   string      `json:"remember_token"`
	DeletedAt       interface{} `json:"deleted_at"`
}
