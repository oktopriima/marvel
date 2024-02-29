package kontoru

import (
	"fmt"
	"github.com/oktopriima/marvel/core/database"
	"gorm.io/gorm"
)

type kon struct {
	db *gorm.DB
}

func (k kon) Lol() {

	fmt.Println("here")
	return
}

type Kon interface {
	Lol()
}

func NewKon(instance database.DBInstance) Kon {
	return &kon{
		db: instance.Database(),
	}
}
