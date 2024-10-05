package users

import (
	"github.com/labstack/echo/v4"
	userRequest "github.com/oktopriima/marvel/app/entity/request/users"
	"github.com/oktopriima/marvel/app/modules/base/response"
	"github.com/oktopriima/marvel/app/usecase/users"
	"github.com/oktopriima/marvel/core/tracer"
	"go.elastic.co/apm/v2"
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
	span, ctx := apm.StartSpan(c.Request().Context(), "userHandler:FindByID", tracer.ProcessTraceName)
	defer span.End()
	var req userRequest.FindByIDRequest

	err := c.Bind(&req)
	if err != nil {
		return response.ErrorResponse(c, err, http.StatusBadRequest)
	}

	output, err := h.userUsecase.FindByID(ctx, req.ID)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SingleResponseData(c, output)
}

func (h UserHandler) FindByEmail(c echo.Context) error {
	span, ctx := apm.StartSpan(c.Request().Context(), "userHandler:FindByEmail", tracer.ProcessTraceName)
	defer span.End()

	var req userRequest.FindByEmailRequest

	err := c.Bind(&req)
	if err != nil {
		return response.ErrorResponse(c, err, http.StatusBadRequest)
	}

	output, err := h.userUsecase.FindByEmail(ctx, req.Email)
	if err != nil {
		return response.ErrorResponse(c, err)
	}
	return response.SingleResponseData(c, output)
}
