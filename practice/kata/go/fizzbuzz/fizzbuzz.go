package fizzbuzz

import "strconv"

// DivisibleByThree returns true if the input is divisible by 3.
func DivisibleByThree(n int) bool {
	return n%3 == 0
}

// DivisibleByFive returns true if the number is divisible by 5.
func DivisibleByFive(n int) bool {
	return n%5 == 0
}

// DivisibleByThreeAndFive returns true if the number is divisible by 3 and 5.
func DivisibleByThreeAndFive(n int) bool {
	return DivisibleByThree(n) && DivisibleByFive(n)
}

// FizzBuzz takes a number and returns "Fizz" if the number is divisible
// by 3, "Buzz" if divisible by 5 and "FizzBuzz" is divisible by both 3 and 5.
// Otherwise, it returns the number as a string.
func FizzBuzz(n int) string {
	if DivisibleByThreeAndFive(n) {
		return "FizzBuzz"
	}
	if DivisibleByThree(n) {
		return "Fizz"
	}
	if DivisibleByFive(n) {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
