package checker

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var formatError = errors.New("invalid email format")
var hostError = errors.New("domain search returned 0 results")
var emailRegexp = regexp.MustCompile(`^\S+@\S+\.\S+$`)

func Format(email string) error {
	if !emailRegexp.MatchString(strings.ToLower(email)) {
		return formatError
	}
	return nil
}

func Host(email string) error {
	if err := Format(email); err != nil {
		return formatError
	}
	s := strings.SplitAfter(email, "@")
	host := s[len(s)-1]

	ips, err := net.LookupIP(host)
	if err != nil {
		return hostError
	}

	if len(ips) <= 0 {
		return hostError
	}
	return nil
}