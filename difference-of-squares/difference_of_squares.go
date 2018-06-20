// Package diffsquares contains the functions to compute the difference of squares
package diffsquares

// SquareOfSums computes the square of the sum of the first _until_ natural numbers
func SquareOfSums(until int) int {
	sum := 0
	for i := 1; i <= until; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares computes the sum of the squares of the first _until_ natural numbers
func SumOfSquares(until int) int {
	sumOfSquares := 0
	for i := 1; i <= until; i++ {
		sumOfSquares += i * i
	}

	return sumOfSquares
}

// Difference returns the difference between the square of sums and the sum of squares
// of the first _until_ natural numbers
func Difference(until int) int {
	return SquareOfSums(until) - SumOfSquares(until)
}
