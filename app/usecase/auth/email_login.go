package auth

import (
	"context"
	"github.com/oktopriima/marvel/app/entity/models"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
	"github.com/oktopriima/marvel/app/usecase/auth/dto/filter"
	"github.com/oktopriima/thor/jwt"
	"strconv"
)

func (a *authenticationUsecase) EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error) {
	var u models.Users

	f := filter.NewLoginFilter(request)

	err := a.userRepo.Search(ctx, &u, f)
	if err != nil {
		return nil, err
	}

	token, err := a.jwtToken.GenerateToken(jwt.Params{
		ID: strconv.Itoa(int(u.Id)),
	})
	if err != nil {
		return nil, err
	}

	return dto.CreateResponse(token), nil
}
