package pkg

import (
	"github.com/go-playground/validator/v10"
)

func ValidateFroms(from interface{}) error {
	validate := validator.New()
	err := validate.Struct(from)
	return err
}
