package implementer

import (
	"github.com/oktopriima/marvel/cmd/seeder/seed"
	"gorm.io/gorm"
)

type SeederImplementer interface {
	seed.SeederInterface
}

func NewSeederImplementer(db *gorm.DB) SeederImplementer {
	return seed.NewBaseSeeder(db)
}
