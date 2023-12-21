package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oktopriima/marvel/app/handler/users"
)

func NewRouter(
	e *echo.Echo,
	userHandler users.UserHandler,
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

	{
		userRoute := route.Group("users")
		userRoute.GET("/:id", userHandler.FindByID)
	}

}
