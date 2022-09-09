package validate

import (
	"errors"

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

func EntStringValidator(tag string) func(s string) error {
	return func(s string) error {
		err := Get().Var(s, tag)
		if err == nil {
			return nil
		}
		var err2 validator.ValidationErrors
		if errors.As(err, &err2) {
			var msg string
			for _, err3 := range err2 {
				msg += err3.Error()
			}
			return errors.New(msg)
		}
		return err
	}
}
func EntEnumValidator(tag string) func(s string) error {
	return func(s string) error {
		err := Get().Var(s, tag)
		if err == nil {
			return nil
		}
		var err2 validator.ValidationErrors
		if errors.As(err, &err2) {
			var msg string
			for _, err3 := range err2 {
				msg += err3.Error()
			}
			return errors.New(msg)
		}
		return err
	}
}