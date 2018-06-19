// Package reverse lets us reverse things
package reverse

import (
	"strings"
)

// String returns a reverse of the input string
func String(input string) string {
	characters := strings.Split(input, "")
	var resultBuilder strings.Builder

	for i := len(characters) - 1; i >= 0; i-- {
		resultBuilder.WriteString(characters[i])
	}

	return resultBuilder.String()
}
