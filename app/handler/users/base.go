package users

import (
	"github.com/labstack/echo/v4"
	userRequest "github.com/oktopriima/marvel/app/entity/request/users"
	"github.com/oktopriima/marvel/app/modules/base/response"
	"github.com/oktopriima/marvel/app/usecase/users"
	"net/http"
)

type UserHandler struct {
	userUsecase users.UserUsecase
}

func NewUserHandler(usecase users.UserUsecase) UserHandler {
	return UserHandler{
		userUsecase: usecase,
	}
}

func (h UserHandler) FindByID(c echo.Context) error {
	var req userRequest.FindByIDRequest

	err := c.Bind(&req)
	if err != nil {
		return response.ErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	output, err := h.userUsecase.FindByID(ctx, req.ID)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SingleResponseData(c, output)
}
