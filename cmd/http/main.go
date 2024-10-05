package main

import (
	"github.com/oktopriima/marvel/cmd"
	"github.com/oktopriima/marvel/cmd/http/router"
	"github.com/oktopriima/marvel/cmd/http/server"
	"os"
)

func main() {
	_ = os.Setenv("ELASTIC_APM_LOG_LEVEL", "debug")

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
