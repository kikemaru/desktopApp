package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	V *validator.Validate
}

func InitValidator() *Validator {
	val := validator.New()

	return &Validator{V: val}
}
