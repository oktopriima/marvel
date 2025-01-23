package dto

import (
	"github.com/oktopriima/thor/jwt"
	"time"
)

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token        string    `json:"token"`
	ExpiredAt    time.Time `json:"expired_at"`
	RefreshToken string    `json:"refresh_token"`
}

func (l *loginResponse) GetObject() *loginResponse {
	return l
}

type LoginResponse interface {
	GetObject() *loginResponse
}

func CreateResponse(response jwt.TokenResponse) LoginResponse {
	return &loginResponse{
		Token:        response.GetStringToken(),
		ExpiredAt:    response.GetTimeExpiredAt(),
		RefreshToken: response.GetStringRefreshToken(),
	}
}
