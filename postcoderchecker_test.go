package checker

import "testing"

func TestPostcodeValidFormat(t *testing.T) {
	if err := Postcode().Format("AA1 1BB"); err != nil {
		t.Errorf("Postcode not valid")
	}
}

func TestPostcodeInvalidFormat(t *testing.T) {
	if err := Postcode().Format("A 1BB"); err == nil {
		t.Errorf("Postcode not valid")
	}
}
