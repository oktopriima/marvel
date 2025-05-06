package users

import (
	"context"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
)

func (u *userUsecase) NotifyLogin(ctx context.Context, req *dto.NotifyLoginRequest) (dto.UserResponse, error) {
	// business logics happen here
	user := new(models.Users)

	user.Id = req.Id
	user.Email = req.Email
	user.Name = req.Name

	return dto.ConvertToResponse(user), nil
}
