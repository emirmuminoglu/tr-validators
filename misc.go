package trvalidators

import (
	"github.com/go-playground/validator/v10"
)

func validateLicensePlate(fl validator.FieldLevel) bool {
	return licensePlateRegex.MatchString(fl.Field().String())
}
