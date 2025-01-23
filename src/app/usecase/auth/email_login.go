package auth

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/src/app/helper"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"github.com/oktopriima/thor/jwt"
	"strconv"
)

func (a *authenticationUsecase) EmailLoginUsecase(ctx context.Context, request dto.EmailLoginRequest) (dto.LoginResponse, error) {
	user, err := a.userRepo.FindByEmail(request.Email, ctx)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(request.Password, user.Password) {
		return nil, fmt.Errorf("password not match")
	}

	token, err := a.jwtToken.GenerateToken(jwt.Params{
		ID: strconv.Itoa(int(user.Id)),
	})
	if err != nil {
		return nil, err
	}

	return dto.CreateResponse(token), nil
}
