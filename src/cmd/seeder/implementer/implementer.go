package implementer

import (
	seed2 "github.com/oktopriima/marvel/src/cmd/seeder/seed"
	"gorm.io/gorm"
)

type SeederImplementer interface {
	seed2.SeederInterface
}

func NewSeederImplementer(db *gorm.DB) SeederImplementer {
	return seed2.NewBaseSeeder(db)
}
