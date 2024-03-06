package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/app/modules/base/response"
	"github.com/oktopriima/marvel/app/usecase/auth"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
	"net/http"
)

type AuthenticationHandler struct {
	uc auth.AuthenticationUsecase
}

func NewAuthenticationHandler(uc auth.AuthenticationUsecase) AuthenticationHandler {
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
