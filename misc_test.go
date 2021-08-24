package trvalidators

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type miscTestStruct struct {
	LicensePlate string `validate:"trlicenseplate"`
}

func TestMisc(t *testing.T) {
	v := validator.New()

	Init(v)

	err := v.Struct(miscTestStruct{
		LicensePlate: "34A2344",
	})

	if err != nil {
		t.Errorf("failed to validate misc: %s", err)
	}
}
