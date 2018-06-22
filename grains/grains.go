// Package grains lets us compute the number of grains that someone ows
package grains

import (
	"errors"
)

// Values is a global cache of square values.
// As more squares are computed, they are added here
var Values = map[int]uint64{
	0: 0,
	1: 1,
	2: 2,
	3: 4,
}

// Square returns the number of grains on a particular chessboard square
func Square(until int) (uint64, error) {
	if until <= 0 {
		return 0, errors.New("Cannot send 0")
	}

	if until > 64 {
		return 0, errors.New("Cannot compute this big a number")
	}

	if Values[until] == 0 {
		val, _ := Square(until - 1)
		Values[until] = val * 2
	}

	return Values[until], nil
}

// Total returns the total number of grains
func Total() uint64 {
	var sum uint64
	for _, value := range Values {
		sum += uint64(value)
	}
	return sum
}
