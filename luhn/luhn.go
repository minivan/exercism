// Package luhn lets us validate numbers usign the Luhn algorithm
package luhn

import (
	"strings"
)

// Valid returns true if the number is valid according to the Luhn algorithm
func Valid(longNumber string) bool {
	// 0. Remove spaces
	longNumber = strings.Replace(longNumber, " ", "", -1)
	if len(longNumber) < 2 {
		return false
	}
	// 1. convert the string to a slice of digits
	var numbers []int

	for _, char := range longNumber {
		digit := int(char - '0')
		if digit < 0 {
			// There was a non-digit in the string. Invalid
			return false
		}
		numbers = append(numbers, digit)
	}

	// 2. double the even digits
	for i := len(numbers) - 2; i >= 0; i -= 2 {
		numbers[i] *= 2
		if numbers[i] > 9 {
			numbers[i] -= 9
		}
	}

	// 3. sum the digits
	sum := 0
	for _, digit := range numbers {
		sum += digit
	}

	// 4. Compute result
	if sum%10 == 0 {
		return true
	}
	return false
}
