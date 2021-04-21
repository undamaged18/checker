package checker

import (
	"errors"
	"regexp"
	"strings"
)

type phone struct {
	Format    func(number string) error
	Mobile    func(number string) bool
	LandLine  func(number string) bool
	Premium   func(number string) bool
	FreePhone func(number string) bool
}

var phoneRegexp = regexp.MustCompile(`^((((\(?0\d{4}\)?\s?\d{3}\s?\d{3})|(\(?0\d{3}\)?\s?\d{3}\s?\d{4})|(\(?0\d{2}\)?\s?\d{4}\s?\d{4}))(\s?\(\d{4}|\d{3}))?)|((\+44\s?7\d{3}|\(?07\d{3}\)?)\s?\d{3}\s?\d{3})|((((\+44\s?\d{4}|\(?0\d{4}\)?)\s?\d{3}\s?\d{3})|((\+44\s?\d{3}|\(?0\d{3}\)?)\s?\d{3}\s?\d{4})|((\+44\s?\d{2}|\(?0\d{2}\)?)\s?\d{4}\s?\d{4}))(\s?\(\d{4}|\d{3}))?$`)
var phoneFormatErr = errors.New("invalid phone number format")

func Phone() *phone {
	return &phone{
		Format:    phoneFormat,
		Mobile:    isMobile,
		LandLine:  isLandline,
		Premium:   isPremiumRate,
		FreePhone: isFreephone,
	}
}

func phoneFormat(number string) error {
	if !phoneRegexp.MatchString(number) {
		return phoneFormatErr
	}
	return nil
}

func isMobile(number string) bool {
	if strings.HasPrefix(number, "07") || strings.HasPrefix(number, "+447") || strings.HasPrefix(number, "+44 7") {
		return true
	}
	return false
}

func isLandline(number string) bool {
	if !strings.HasPrefix(number, "07") || !strings.HasPrefix(number, "+447") || !strings.HasPrefix(number, "+44 7") {
		return true
	}
	return false
}

func isPremiumRate(number string) bool {
	prefixes := []string{
		"084",
		"+4484",
		"+44 84",
		"087",
		"+4487",
		"+44 87",
		"09",
		"+449",
		"+44 9",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}
	return false
}

func isFreephone(number string) bool {
	prefixes := []string{
		"0800",
		"+44800",
		"+44 800",
		"0808",
		"+44808",
		"+44 808",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}
	return false
}
