package crypto

import (
	"strings"

	"github.com/matribots/go-swiss-knife/charstring"
)

type ascii byte

// Generate a random strong password with given length
func RandomPassword(len int) *string {
	asciiChars := make([]ascii, len)
	password := string(asciiChars)
	for !isValidPassword(&password) {
		password = string(charstring.RandomPrintableASCII(33, 127, len))
	}
	return &password
}

// Check if a password contains special characters such as ":", "?", "*" and numbers
func isValidPassword(password *string) bool {
	// A very ungraceful implementation. Will rewrite later.
	if (strings.ContainsRune(*password, ':') ||
		strings.ContainsRune(*password, '?') ||
		strings.ContainsRune(*password, '*') ||
		strings.ContainsRune(*password, '+') ||
		strings.ContainsRune(*password, '-') ||
		strings.ContainsRune(*password, '=') ||
		strings.ContainsRune(*password, ':') ||
		strings.ContainsRune(*password, ';') ||
		strings.ContainsRune(*password, '$')) &&
		(strings.ContainsRune(*password, '0') ||
			strings.ContainsRune(*password, '1') ||
			strings.ContainsRune(*password, '2') ||
			strings.ContainsRune(*password, '3') ||
			strings.ContainsRune(*password, '4') ||
			strings.ContainsRune(*password, '5') ||
			strings.ContainsRune(*password, '6') ||
			strings.ContainsRune(*password, '7') ||
			strings.ContainsRune(*password, '8') ||
			strings.ContainsRune(*password, '9')) {
		return true
	} else {
		return false
	}
}
