package users

import (
	"context"
	"github.com/oktopriima/marvel/pkg/tracer"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
	"go.elastic.co/apm/v2"
)

func (u *userUsecase) FindByEmail(ctx context.Context, email string) (output *dto.UserResponse, err error) {
	span, ctx := apm.StartSpan(ctx, "userUsecase.FindByEmail", tracer.ProcessTraceName)
	defer span.End()

	user, err := u.userRepo.FindByEmail(email, ctx)
	if err != nil {
		return nil, err
	}

	output = new(dto.UserResponse)
	return output.ConvertToResponse(user), nil
}
