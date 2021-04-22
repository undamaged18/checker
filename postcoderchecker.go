package checker

import "regexp"

type postcode struct {
	Format func(postcode string) bool
}

var postcodeRegex = regexp.MustCompile(`^([A-Z]{1,2}\d[A-Z\d]? ?\d[A-Z]{2}|GIR ?0A{2})$`)

func Postcode() *postcode {
	return &postcode{Format: postcodeFormat}
}

func postcodeFormat(postcode string) bool {
	return postcodeRegex.MatchString(postcode)
}
