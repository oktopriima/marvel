package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/src/app/modules/base/response"
	"github.com/oktopriima/marvel/src/app/usecase/auth"
	"github.com/oktopriima/marvel/src/app/usecase/auth/dto"
	"net/http"
)

type AuthenticationHandler struct {
	uc auth.AuthenticationUsecaseContract
}

func NewAuthenticationHandler(uc auth.AuthenticationUsecaseContract) AuthenticationHandler {
	return AuthenticationHandler{
		uc: uc,
	}
}

func (a AuthenticationHandler) LoginByEmail(c echo.Context) error {
	var req dto.EmailLoginRequest

	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	output, err := a.uc.EmailLoginUsecase(ctx, req)
	if err != nil {
		return response.ErrorResponse(c, err, http.StatusForbidden)
	}

	return response.SingleResponseData(c, output)
}
