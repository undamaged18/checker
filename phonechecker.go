package checker

import (
	"errors"
	"github.com/ttacon/libphonenumber"
	"strings"
)

type phone struct {
	Format    func(number string, region string, f format) (string, error)
	Mobile    func(number string) bool
	LandLine  func(number string) bool
	Premium   func(number string) bool
	FreePhone func(number string) bool
}

var phoneFormatErr = errors.New("invalid phone number format")

type format libphonenumber.PhoneNumberFormat

const NATIONAL = 2
const INTERNATIONAL = 1

func Phone() *phone {
	return &phone{
		Format:    phoneFormat,
		Mobile:    isMobile,
		LandLine:  isLandline,
		Premium:   isPremiumRate,
		FreePhone: isFreephone,
	}
}

func phoneFormat(number string, region string, f format) (string, error) {
	num, err := libphonenumber.Parse(number, region)
	if err != nil {
		return "", phoneFormatErr
	}
	var phoneNumber string
	switch f {
	case NATIONAL:
		phoneNumber = libphonenumber.Format(num, libphonenumber.NATIONAL)
	case INTERNATIONAL:
		phoneNumber = libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	default:
		phoneNumber = libphonenumber.Format(num, libphonenumber.NATIONAL)
	}
	return phoneNumber, nil
}

func isMobile(number string) bool {
	return strings.HasPrefix(number, "07") || strings.HasPrefix(number, "+447") || strings.HasPrefix(number, "+44 7")
}

func isLandline(number string) bool {
	return !strings.HasPrefix(number, "07") || !strings.HasPrefix(number, "+447") || !strings.HasPrefix(number, "+44 7")
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
