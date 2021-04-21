package checker

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var FormatErr = errors.New("invalid email format")
var HostErr = errors.New("domain search returned 0 results")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type emailChecker struct {
	Format func(email string) error
	Host   func(email string) error
}

func Email() *emailChecker {
	return &emailChecker{
		Format: emailFormat,
		Host:   emailHost,
	}
}

func emailFormat(email string) error {
	if !emailRegexp.MatchString(strings.ToLower(email)) {
		return FormatErr
	}
	return nil
}

func emailHost(email string) error {
	if err := emailFormat(email); err != nil {
		return FormatErr
	}
	s := strings.SplitAfter(email, "@")
	host := s[len(s)-1]

	ips, err := net.LookupMX(host)
	if err != nil {
		return HostErr
	}

	if len(ips) <= 0 {
		return HostErr
	}
	return nil
}
