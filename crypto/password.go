package crypto

import (
	"regexp"

	"github.com/matribots/go-swiss-knife/charstring"
)

type ascii byte

var numRegExp, lowerCaseRegExp, upperCaseRegExp, symbolRexExp *regexp.Regexp

func init() {
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	numRegExp = regexp.MustCompile(num)
	lowerCaseRegExp = regexp.MustCompile(a_z)
	upperCaseRegExp = regexp.MustCompile(A_Z)
	symbolRexExp = regexp.MustCompile(symbol)
}

// Generate a random strong password with given length
func RandomPassword(len int) *string {
	asciiChars := make([]ascii, len)
	password := string(asciiChars)
	for !isValidPassword(&password) {
		password = string(charstring.RandomPrintableASCII(33, 127, len))
	}
	return &password
}

// Check if a password contains special characters such as ":", "?", "*" and numbers, lowercases characters, uppercase characters
func isValidPassword(password *string) bool {
	// Refer: https://blog.csdn.net/qq_43490063/article/details/104249988
	if !numRegExp.MatchString(*password) {
		return false
	}

	if !lowerCaseRegExp.MatchString(*password) {
		return false
	}

	if !upperCaseRegExp.MatchString(*password) {
		return false
	}

	if !symbolRexExp.MatchString(*password) {
		return false
	}

	return true
}
