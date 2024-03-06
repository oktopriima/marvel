package auth

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
	"github.com/oktopriima/marvel/app/usecase/auth/dto/filter"
)

func (a *authenticationUsecase) EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (*dto.LoginResponse, error) {
	var u models.Users

	f := filter.NewLoginFilter(request)

	err := a.userRepo.Search(ctx, &u, f)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
