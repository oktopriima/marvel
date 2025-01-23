package main

import (
	"github.com/oktopriima/marvel/src"
	"github.com/oktopriima/marvel/src/kafka/consumer/group"
)

func main() {
	c := src.NewRegistry()

	if err := c.Invoke(func(c *group.ConsumerConfig) {
		c.Serve()
	}); err != nil {
		panic(err)
	}
}
