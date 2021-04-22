package checker

import "testing"

func TestPostcodeValidFormat(t *testing.T) {
	if !Postcode().Format("AA1 1BB") {
		t.Errorf("Postcode not valid")
	}
}

func TestPostcodeInvalidFormat(t *testing.T) {
	if Postcode().Format("AA 1 1BB") {
		t.Errorf("Postcode not valid")
	}
}
