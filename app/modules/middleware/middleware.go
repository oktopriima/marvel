package middleware

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oktopriima/marvel/app/helper"
	"github.com/oktopriima/marvel/app/modules/base/response"
	"net/http"
)

const (
	Token    = "TOKEN"
	AuthUser = "AUTH_USER"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			oldCtx := c.Request().Context()

			headerToken, err := helper.HeaderExtractor("Authorization", r)
			if err != nil {
				return response.ErrorResponse(c, err, http.StatusForbidden)
			}

			ctx := context.WithValue(oldCtx, Token, headerToken)
			//ctx = context.WithValue(ctx, AuthUser, e.UserId)

			c.Request().WithContext(ctx)
			return next(c)
		}
	}
}
