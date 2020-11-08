package utils

import (
	"regexp"
	"strings"
)

func IsValidPhoneNumber(val string) bool {
	val = strings.TrimSpace(val)
	reg := regexp.MustCompile(`^1[0-9]{10}$`)
	return reg.MatchString(val)
}
