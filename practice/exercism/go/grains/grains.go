// Package grains calculates the grains on a chess board starting with 1 grain on the first square,
// then doubling on each square thereafter.
package grains

import "github.com/pkg/errors"

// Square returns the number of grains on a particular Square of a chessboard
func Square(n int) (uint64, error) {

	if n < 1 || n > 64 {
		return 0, errors.New("Specify a square between 1 and 64")
	}

	return pow(2, uint64(n-1)), nil
}

// Total calculates the total number of grains on a chessboard
func Total() uint64 {

	grains := uint64(1)
	for n := uint64(1); n <= 64; n++ {
		grains += pow(2, n)
	}

	return grains
}

// pow uses exponentiation by squaring to calculate x to the power of n, faster than multiplying by itself n times.
// See here: https://simple.wikipedia.org/wiki/Exponentiation_by_squaring
func pow(x, n uint64) uint64 {

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n%2 == 0 {
		return pow(x*x, n/2)
	}

	return x * pow(x*x, (n-1)/2)
}
