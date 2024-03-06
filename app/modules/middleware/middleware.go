package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/app/helper"
	"github.com/oktopriima/marvel/app/modules/base/response"
	"github.com/oktopriima/thor/jwt"
	"net/http"
)

const (
	Token    = "TOKEN"
	AuthUser = "AUTH_USER"
)

func Auth(token jwt.AccessToken) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			oldCtx := c.Request().Context()

			headerToken, err := helper.HeaderExtractor("Authorization", r)
			if err != nil {
				return response.ErrorResponse(c, err, http.StatusForbidden)
			}

			if !token.Validate(headerToken) {
				return response.ErrorResponse(c, fmt.Errorf("token invalid"), http.StatusForbidden)
			}

			e, err := jwt.Extract(headerToken, token.GetSignatureKey())
			if err != nil {
				return response.ErrorResponse(c, err, http.StatusForbidden)
			}

			ctx := context.WithValue(oldCtx, Token, headerToken)
			ctx = context.WithValue(ctx, AuthUser, e.Id)

			c.Request().WithContext(ctx)
			return next(c)
		}
	}
}
