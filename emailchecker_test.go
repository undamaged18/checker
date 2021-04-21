package checker

import (
	"testing"
)

func TestCorrectFormat(t *testing.T) {
	var email = "valid@example.com"
	err := Email().Format(email)
	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestIncorrectFormatNoDomain(t *testing.T) {
	var email = "invalid@"
	err := Email().Format(email)
	if err != nil {
		if err != FormatErr {
			t.Errorf("Email Format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email Format failed: should of returned error as")
	}
}

func TestValidHost(t *testing.T) {
	var email = "valid@gmail.com"
	err := Email().Host(email)

	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestInvalidHost(t *testing.T) {
	var email = "invalid@invalid_hostname.com"
	err := Email().Host(email)

	if err != nil {
		if err != FormatErr && err != HostErr {
			t.Errorf("Email Host validation failed: %v", err.Error())
		}
	}
}
