package users

import (
	"context"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/app/entity/models"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
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
