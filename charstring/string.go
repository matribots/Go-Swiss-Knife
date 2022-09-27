package charstring

import "strings"

// Check if a given string is empty.
// Both "" and "   " will be regonized as an empty string.
func IsEmptyString(str *string) bool {
	if str == nil {
		panic("The param \"str\" cannot be nil!")
	}
	if len(*str) == 0 || strings.TrimSpace(*str) == "" {
		return true
	} else {
		return false
	}
}
