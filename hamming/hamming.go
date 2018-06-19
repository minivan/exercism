// Package hamming computes the hamming distance between two strings
package hamming

import "errors"

// Distance takes two strings and returns the hamming distance
func Distance(a, b string) (int, error) {
	acc := 0
	if len(a) != len(b) {
		return 0, errors.New("The string length should be the same")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			acc++
		}
	}
	return acc, nil
}
