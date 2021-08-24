package trvalidators

import "regexp"

var (
	licensePlateRegex  = regexp.MustCompile(`^(0[1-9]|[1-7][0-9]|8[01])(([A-Z])(\d{4,5})|([A-Z]{2})(\d{3,4})|([A-Z]{3})(\d{2,3}))$`)
	landPhoneRegex     = regexp.MustCompile(`^(0)([2348]{1})([0-9]{2})\s?([0-9]{3})\s?([0-9]{2})\s?([0-9]{2})$`)
	cellPhoneRegex     = regexp.MustCompile(`^(05)([0-9]{2})\s?([0-9]{3})\s?([0-9]{2})\s?([0-9]{2})$`)
	citizenshipIDRegex = regexp.MustCompile(`^[1-9]{1}[0-9]{9}[02468]{1}$`)
	taxIDRegex         = regexp.MustCompile(`^[0-9]{10}$`)
)
