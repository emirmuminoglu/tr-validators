package trvalidators

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type phonenumberTestStruct struct {
	CellNumber string `validate:"trcellphone"`
	LandNumber string `validate:"trlandphone"`
}

func TestPhoneNumbers(t *testing.T) {
	v := validator.New()

	Init(v)

	err := v.Struct(phonenumberTestStruct{
		CellNumber: "05231231212",
		LandNumber: "02231231212",
	})

	if err != nil {
		t.Errorf("failed to validate misc: %s", err)
	}
}
