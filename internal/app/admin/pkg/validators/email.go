package validators

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func IsEmailValid(email string) bool {
	return emailRegex.MatchString(strings.ToLower(email))
}
