// Package isogram can help with guessing which words are isograms
package isogram

import (
	"fmt"
	"strings"
)

// IsIsogram takes a string and returns true if it is an isogram
func IsIsogram(input string) bool {
	characterCounts := make(map[string]int)
	byteset := []byte(input)

	for _, character := range byteset {
		strChar := strings.ToLower(string(character))
		if characterCounts[strChar] != 0 &&
			strChar != " " &&
			strChar != "-" {
			return false
		}
		fmt.Println(characterCounts)
		characterCounts[strChar] = 1
	}

	return true
}
