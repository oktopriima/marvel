package validates

import (
	"github.com/go-playground/validator/v10"
)

type validatorLib struct {
	Valid *validator.Validate
}

// New ...
func New() Validator {
	return &validatorLib{
		Valid: validator.New(),
	}
}

func (v *validatorLib) Request(value interface{}) error {
	return v.Valid.Struct(value)
}
