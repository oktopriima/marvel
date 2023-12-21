package response

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type SingleApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MultipleApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

func SingleResponseData(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, SingleApiResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("OK"),
		Data:    data,
	})
}

func MultipleResponseData(ctx echo.Context, data interface{}, page, size int) error {
	return ctx.JSON(http.StatusOK, MultipleApiResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("OK"),
		Data:    data,
		Page:    page,
		Size:    size,
	})
}

func ErrorResponse(ctx echo.Context, err error, statusCode ...int) error {
	code := http.StatusUnprocessableEntity
	if len(statusCode) >= 0 {
		code = statusCode[0]
	}
	return ctx.JSON(code, ErrorApiResponse{
		Code:    code,
		Message: err.Error(),
	})
}
