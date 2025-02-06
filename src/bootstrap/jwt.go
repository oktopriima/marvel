package bootstrap

import (
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/thor/jwt"
	"go.uber.org/dig"
)

func NewJWT(container *dig.Container) *dig.Container {
	var err error
	// provide auth token
	if err = container.Provide(func(cfg config.AppConfig) jwt.AccessToken {
		return jwt.NewAccessToken(jwt.Request{
			SignatureKey: cfg.Jwt.Key,
			Audience:     "",
			Issuer:       cfg.Jwt.Issuer,
		})
	}); err != nil {
		panic(err)
	}

	return container
}
