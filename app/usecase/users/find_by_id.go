package users

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
	"github.com/oktopriima/marvel/core/tracer"
	"go.elastic.co/apm/v2"
)

func (u *userUsecase) FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error) {
	span, ctx := apm.StartSpan(ctx, "userUsecase.FindByID", tracer.ProcessTraceName)
	defer span.End()
	m := new(models.Users)
	err := u.userRepo.FindByID(ctx, m, ID)
	if err != nil {
		return nil, err
	}

	output := new(dto.UserResponse)

	return output.ConvertToResponse(m), nil
}
