package main

import (
	"github.com/oktopriima/marvel/cmd"
	"github.com/oktopriima/marvel/cmd/http/router"
	"github.com/oktopriima/marvel/cmd/http/server"
)

func main() {
	c := cmd.NewRegistry()

	err := c.Invoke(router.NewRouter)
	if err != nil {
		panic(err)
	}

	if err := c.Invoke(func(instance *server.EchoInstance) {
		instance.RunWithGracefullyShutdown()
	}); err != nil {
		panic(err)
	}
}
