package checker

import (
	"testing"
)

func TestPhoneValidFormatNational(t *testing.T) {
	num, err := Phone().Format("+44 7777 555 555", "GB", NATIONAL)
	if err != nil {
		t.Errorf("Phone Format failed: %v", err.Error())
	}
	if num == "" {
		t.Errorf("Phone Format failed: %v", "Expected returned number to be in format +44 xxxx xxxxxx")
	}
}

func TestPhoneValidFormatInternational(t *testing.T) {
	num, err := Phone().Format("+44 7777 555 555", "GB", INTERNATIONAL)
	if err != nil {
		t.Errorf("Phone Format failed: %v", err.Error())
	}
	if num == "" {
		t.Errorf("Phone Format failed: %v", "Expected returned number to be in format +44 xxxx xxxxxx")
	}
}

func TestPhoneInvalidFormat(t *testing.T) {
	_, err := Phone().Format("This is not a phone number", "GB", INTERNATIONAL)
	if err == nil {
		t.Errorf("Phone Format failed: %v", "Expected a format failure")
	}
}

func TestPhoneMobile(t *testing.T) {
	if !Phone().Mobile("+44 7777 555 555") {
		t.Errorf("Phone Mobile check failed")
	}
}

func TestPhoneNonMobile(t *testing.T) {
	if Phone().Mobile("+44 1234 555 555") {
		t.Errorf("Phone Non Mobile check failed")
	}
}

func TestPhonePremium(t *testing.T) {
	if !Phone().Premium("+44 845 0000 000") {
		t.Errorf("Phone Premium check failed")
	}
}

func TestPhoneNonPremium(t *testing.T) {
	if Phone().Premium("+44 234 0000 000") {
		t.Errorf("Phone Non Premium check failed")
	}
}

func TestPhoneFreePhone(t *testing.T) {
	if !Phone().FreePhone("+44 800 0000 000") {
		t.Errorf("Phone Freephone check failed")
	}
}

func TestPhoneNonFreePhone(t *testing.T) {
	if Phone().FreePhone("+44 234 0000 000") {
		t.Errorf("Phone Freephone non check failed")
	}
}
