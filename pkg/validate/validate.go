package validate

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

func Get() *validator.Validate {
	if Validate == nil {
		Validate = validator.New()
	}
	return Validate
}
