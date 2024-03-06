package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/marvel/core/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	EchoInstance struct {
		Router *echo.Echo
		Config config.AppConfig
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func NewEchoInstance(r *echo.Echo, cfg config.AppConfig) *EchoInstance {
	return &EchoInstance{
		Router: r,
		Config: cfg,
	}
}

func (server *EchoInstance) runHttp() (err error) {
	port := fmt.Sprintf(":%s", server.Config.App.Port)
	server.Router.Validator = &CustomValidator{validator: validator.New()}
	if err = server.Router.Start(port); err != nil {
		return err
	}

	return
}

func (server *EchoInstance) RunWithGracefullyShutdown() {
	// run server on another thread
	go func() {
		err := server.runHttp()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Router.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return err
	}

	return nil
}
