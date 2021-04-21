package checker

import "testing"

func TestCorrectFormat(t *testing.T) {
	var email = "valid.1_2-3@e_x-a.mple.com"
	err := Format(email)
	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestIncorrectFormatNoDomain(t *testing.T) {
	var email = "invalid@"
	err := Format(email)
	if err != nil {
		if err.Error() != "invalid email format" {
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
		if err.Error() != "invalid email format" {
			t.Errorf("Email Format failed: %v", err.Error())
		}
	}
	if err == nil {
		t.Errorf("Email Format failed: should of returned error as")
	}
}

func TestValidHost(t *testing.T) {
	var email = "valid@blueyonder.com"
	err := Host(email)

	if err != nil {
		t.Errorf("Email Format failed: %v", err.Error())
	}
}

func TestInvalidHost(t *testing.T) {
	var email = "invalid@invalid_hostname_should_not_return_value.com"
	err := Host(email)

	if err != nil {
		if err.Error() != "domain search returned 0 results" {
			t.Errorf("Email host validation failed: %v", err.Error())
		}
	}
}