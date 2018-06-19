// Package bob contains all the functions to simulate Bob the teenager
package bob

import "strings"

// Sanitize removes extra whitespace from the input string
func Sanitize(str string) string {
	return strings.TrimSpace(str)
}

// IsUpper returns true if the input is all upppercase characters
func IsUpper(str string) bool {
	return str == strings.ToUpper(str) && str != strings.ToLower(str)
}

// IsBlank checks whether the input is blank
func IsBlank(str string) bool {
	return Sanitize(str) == ""
}

// IsQuestion checks whether the input is a question
func IsQuestion(str string) bool {
	return strings.HasSuffix(Sanitize(str), "?")
}

// IsYelledSentence checks whether the input is not a question and is yelled
func IsYelledSentence(str string) bool {
	return IsUpper(str) && !IsQuestion(str)
}

// IsYelledQuestion returns if the input is both a question and is yelled
func IsYelledQuestion(str string) bool {
	return IsUpper(str) && IsQuestion(str)
}

// Hey simulates Bob, the lackadaisical teenager
func Hey(remark string) string {
	switch {
	case IsYelledSentence(remark):
		return "Whoa, chill out!"
	case IsYelledQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case IsQuestion(remark):
		return "Sure."
	case IsBlank(remark):
		return "Fine. Be that way!"
	}
	return "Whatever."
}
