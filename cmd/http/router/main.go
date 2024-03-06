package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oktopriima/marvel/app/handler/auth"
	"github.com/oktopriima/marvel/app/handler/users"
	jwtMidl "github.com/oktopriima/marvel/app/modules/middleware"
	"github.com/oktopriima/thor/jwt"
)

func NewRouter(
	e *echo.Echo,
	jwtAuth jwt.AccessToken,
	userHandler users.UserHandler,
	authHandler auth.AuthenticationHandler,
) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	route := e.Group("api/")

	route.GET("ping", func(context echo.Context) error {
		return context.JSON(200, struct {
			Data interface{} `json:"data"`
		}{Data: "pong"})
	})

	// login route
	{
		loginRoute := route.Group("auth")
		loginRoute.POST("/email", authHandler.LoginByEmail)
	}

	// authenticate route
	{
		meRoute := route.Group("me")
		meRoute.Use(jwtMidl.Auth(jwtAuth))
		meRoute.GET("", func(c echo.Context) error {
			return c.JSON(200, struct {
				Message string `json:"message"`
			}{Message: "OK"})
		})
	}

	{
		userRoute := route.Group("users")
		userRoute.GET("/:id", userHandler.FindByID)
	}

}
