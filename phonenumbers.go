package trvalidators

import (
	"github.com/go-playground/validator/v10"
)

func validateLandPhone(fl validator.FieldLevel) bool {
	return landPhoneRegex.MatchString(fl.Field().String())
}

func validateCellPHone(fl validator.FieldLevel) bool {
	return cellPhoneRegex.MatchString(fl.Field().String())
}
