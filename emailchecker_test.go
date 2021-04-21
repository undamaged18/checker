package checker

import (
	"testing"
)

func TestCorrectFormat(t *testing.T) {
	var email = "valid@example.com"
	err := Format(email)
	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestIncorrectFormatNoDomain(t *testing.T) {
	var email = "invalid@"
	err := Format(email)
	if err != nil {
		if err != formatError {
			t.Errorf("Email Format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email Format failed: should of returned error as")
	}
}

func TestIncorrectFormatNoTLD(t *testing.T) {
	var email = "invalid@example"
	err := Format(email)
	if err != nil {
		if err != formatError {
			t.Errorf("Email Format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email Format failed: should of returned error as: %s", formatError)
	}
}

func TestValidHost(t *testing.T) {
	var email = "valid@gmail.com"
	err := Host(email)

	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestInvalidHost(t *testing.T) {
	var email = "invalid@invalid_hostname.com"
	err := Host(email)

	if err != nil {
		if err != formatError && err != hostError {
			t.Errorf("Email host validation failed: %v", err.Error())
		}
	}
}
