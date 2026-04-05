package utils

import "strings"

func CapFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}