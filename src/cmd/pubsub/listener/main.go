package main

import (
	"context"
	"github.com/oktopriima/marvel/src/bootstrap/pubsub"
	"github.com/oktopriima/marvel/src/cmd/pubsub/listener/handler"
	"github.com/oktopriima/marvel/src/cmd/pubsub/listener/router"
	"github.com/oktopriima/marvel/src/cmd/pubsub/listener/server"
)

func main() {
	var err error

	c := pubsub.Bootstrap()

	if err = c.Provide(router.NewRouter); err != nil {
		panic(err)
	}

	if err := c.Provide(router.NewRegisteredDomain); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(r handler.Router) {
		server.NewServer(r).Run(context.Background())
	}); err != nil {
		panic(err)
	}
}
