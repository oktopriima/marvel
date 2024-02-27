package main

import (
	"fmt"
	"github.com/oktopriima/marvel/cmd"
	"log"
)

func main() {
	c := cmd.NewRegistry()

	err := c.Invoke(migrate())
	if err != nil {
		log.Fatalf("error while run migration. message : %v", err)
	}

	log.Printf("SUCCESS")
}

func migrate() error {
	fmt.Println("here")
	return nil
}
