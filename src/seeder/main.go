package main

import (
	"context"
	"github.com/oktopriima/marvel/pkg/database"
	"github.com/oktopriima/marvel/src"
	"github.com/oktopriima/marvel/src/seeder/seed/users"
	"gorm.io/gorm"
)

func main() {
	c := src.NewRegistry()

	// provide the seeder interface
	if err := c.Provide(NewSeeder); err != nil {
		panic(err)
	}

	err := c.Invoke(func(s Seeder) {
		s.Init()
	})
	if err != nil {
		panic(err)
	}

}

type seeder struct {
	db  *gorm.DB
	ctx context.Context
}

type Seeder interface {
	Init()
}

func NewSeeder(instance database.DBInstance) Seeder {
	return &seeder{
		db:  instance.Database(),
		ctx: context.Background(),
	}
}

func (s *seeder) Init() {
	// register the seeder in here
	users.Run(s.ctx, s.db)

	return
}
