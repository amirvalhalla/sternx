package util

import "strings"

func IsStringEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}
