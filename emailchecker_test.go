package checker

import (
	"testing"
)

func TestCorrectFormat(t *testing.T) {
	var email = "valid@example.com"
	err := Email().Format(email)
	if err != nil {
		t.Errorf("Email format failed: %v", err.Error())
	}
}

func TestIncorrectFormatNoDomain(t *testing.T) {
	var email = "invalid@"
	err := Email().Format(email)
	if err != nil {
		if err != FormatErr {
			t.Errorf("Email format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email format failed: should of returned error as")
	}
}

func TestIncorrectFormatNoTLD(t *testing.T) {
	var email = "invalid@example"
	err := Email().Format(email)
	if err != nil {
		if err != FormatErr {
			t.Errorf("Email format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email format failed: should of returned error as: %s", FormatErr)
	}
}

func TestValidHost(t *testing.T) {
	var email = "valid@gmail.com"
	err := Email().Host(email)

	if err != nil {
		t.Errorf("Email format failed: %v", err.Error())
	}
}

func TestInvalidHost(t *testing.T) {
	var email = "invalid@invalid_hostname.com"
	err := Email().Host(email)

	if err != nil {
		if err != FormatErr && err != HostErr {
			t.Errorf("Email host validation failed: %v", err.Error())
		}
	}
}
