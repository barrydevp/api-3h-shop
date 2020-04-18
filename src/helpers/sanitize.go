package helpers

import "strings"

func SanitizeString(s string) string {
	return strings.Trim(s, " ")
}
