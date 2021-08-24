package trvalidators

import "github.com/go-playground/validator/v10"

func Init(v *validator.Validate) {
	v.RegisterValidation("trcellphone", validateCellPHone)
	v.RegisterValidation("trlandphone", validateLandPhone)
	v.RegisterValidation("trtaxid", validateTaxID)
	v.RegisterValidation("trcitizenshipid", validateCitizenshipID)
	v.RegisterValidation("trlicenseplate", validateLicensePlate)
}
