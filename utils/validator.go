package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Validator struct{}

func (Validator) ValidateBody(data interface{}) []string {
	var errors []string
	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("%s is a %s field", err.Field(), err.Tag())
			errors = append(errors, msg)
		}
	}
	return errors
}
