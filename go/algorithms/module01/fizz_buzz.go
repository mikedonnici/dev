package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var xs []string
	for i := 1; i <= n; i++ {
		xs = append(xs, fizzBuzzOrNumber(i))
	}
	fmt.Printf("%s\n", strings.Join(xs, ", "))
}

func fizzBuzzOrNumber(n int) string {
	switch {
	case n%5 == 0 && n%3 == 0:
		return "Fizz Buzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(n)
}
