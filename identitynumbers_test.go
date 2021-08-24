package trvalidators

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type identityNumbersTestStruct struct {
	TaxID         string `validate:"trtaxid"`
	CitizenshipID string `validate:"trcitizenshipid"`
}

func TestIdentities(t *testing.T) {
	v := validator.New()

	Init(v)

	err := v.Struct(identityNumbersTestStruct{
		TaxID:         "1234567890",
		CitizenshipID: "12345678902",
	})

	if err != nil {
		t.Errorf("failed to validate identities: %s", err)
	}
}
