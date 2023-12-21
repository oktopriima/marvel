package users

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
)

func (u *userUsecase) FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error) {
	m := new(models.Users)
	err := u.userRepo.FindByID(ctx, m, ID)
	if err != nil {
		return nil, err
	}

	output := new(dto.UserResponse)
	output = output.ConvertToResponse(m)

	return output, nil
}
