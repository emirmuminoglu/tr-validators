package trvalidators

import (
	"github.com/go-playground/validator/v10"
)

func validateCitizenshipID(fl validator.FieldLevel) bool {
	return citizenshipIDRegex.MatchString(fl.Field().String())
}

func validateTaxID(fl validator.FieldLevel) bool {
	return taxIDRegex.MatchString(fl.Field().String())
}
