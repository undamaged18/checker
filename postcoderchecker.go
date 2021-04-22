package checker

import (
	"errors"
	"regexp"
)

type postcode struct {
	Format func(postcode string) error
}

var postcodeRegex = regexp.MustCompile(`^([A-Z]{1,2}\d[A-Z\d]? ?\d[A-Z]{2}|GIR ?0A{2})$`)
var postcodeFormatErr = errors.New("invalid UK postcode format")

func Postcode() *postcode {
	return &postcode{Format: postcodeFormat}
}

func postcodeFormat(postcode string) error {
	if postcodeRegex.MatchString(postcode) {
		return nil
	}
	return postcodeFormatErr
}
