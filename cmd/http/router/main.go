package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oktopriima/marvel/app/handler/auth"
	"github.com/oktopriima/marvel/app/handler/users"
	jwtMidl "github.com/oktopriima/marvel/app/modules/middleware"
	"github.com/oktopriima/thor/jwt"
	"go.elastic.co/apm/module/apmechov4/v2"
	"go.elastic.co/apm/v2"
)

func NewRouter(
	e *echo.Echo,
	t *apm.Tracer,
	jwtAuth jwt.AccessToken,
	userHandler users.UserHandler,
	authHandler auth.AuthenticationHandler,
) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(apmechov4.Middleware(apmechov4.WithTracer(t)))

	route := e.Group("api/v1/")

	route.GET("ping", func(context echo.Context) error {
		return context.JSON(200, struct {
			Data interface{} `json:"data"`
		}{Data: "pong"})
	})

	// login route
	{
		loginRoute := route.Group("login")
		loginRoute.POST("/", authHandler.LoginByEmail)
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
		userRoute.GET("/email/:email", userHandler.FindByEmail)
	}

}
