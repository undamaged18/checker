package checker

import (
	"testing"
)

func TestPhoneFormat(t *testing.T) {
	err := Phone().Format("This is not a phone number")
	if err != nil {
		t.Errorf("Phone Format failed: %v", err.Error())
	}
}

func TestPhoneMobile(t *testing.T) {
	if !Phone().Mobile("+44 7777 555 555") {
		t.Errorf("Phone Mobile check failed")
	}
}

func TestPhonePremium(t *testing.T) {
	if !Phone().Premium("+44 845 0000 000") {
		t.Errorf("Phone Premium check failed")
	}
}

func TestPhoneFreePhone(t *testing.T) {
	if !Phone().FreePhone("+44 800 0000 000") {
		t.Errorf("Phone Freephone check failed")
	}
}
