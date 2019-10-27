package fibonacci_test

import (
	"testing"

	"github.com/mikedonnici/dev/go/functional-programming/fibonacci"
)

func TestFibonacciRecursive(t *testing.T) {

	cases := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{20, 6765},
		{42, 267914296},
	}

	for _, c := range cases {
		got := fibonacci.FibonacciRecursive(c.arg)
		if got != c.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", c.arg, got, c.want)
		}
	}
}
