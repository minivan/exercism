// Package scrabble lets us compute scrabble scores for different letter configurations
package scrabble

import (
	"strings"
)

// Score returns the scrabble score of a string
func Score(str string) int {
	values := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	score := 0

	for _, char := range str {
		for letterScore, letterValues := range values {
			uppercaseChar := strings.ToUpper(string(char))

			if strings.Contains(letterValues, uppercaseChar) {
				score += letterScore
			}
		}
	}

	return score
}
