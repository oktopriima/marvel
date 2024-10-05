package users

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/usecase/users/dto"
	"github.com/oktopriima/marvel/core/tracer"
)

func (u *userUsecase) FindByID(ctx context.Context, ID int64) (*dto.UserResponse, error) {
	var (
		err      error
		response *dto.UserResponse
	)

	trace := tracer.StartTrace(ctx, "users:usecase:findById", tracer.UsecaseTraceName)
	defer func() {
		trace.Finish(map[string]interface{}{
			"error":  err,
			"output": response,
		})
	}()

	m := new(models.Users)
	err = u.userRepo.FindByID(trace.Context(), m, ID)
	if err != nil {
		return nil, err
	}

	output := new(dto.UserResponse)
	response = output.ConvertToResponse(m)

	return response, nil
}
