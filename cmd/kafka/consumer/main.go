package main

import (
	"github.com/oktopriima/marvel/cmd"
	"github.com/oktopriima/marvel/cmd/kafka/consumer/group"
)

func main() {
	c := cmd.NewRegistry()

	if err := c.Invoke(func(c *group.ConsumerConfig) {
		c.Serve()
	}); err != nil {
		panic(err)
	}
}
