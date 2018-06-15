// Package luhn implements the luhn algorithm
package luhn

import (
	"log"
	"strconv"
	"strings"
)

// Valid applies the luhn algorithm to a string of numbers, and returns true if valid
func Valid(number string) bool {

	// remove all white space
	number = strings.Replace(number, " ", "", -1)

	// single digit not valid
	if len(number) <= 1 {
		return false
	}

	// only number runes allowed
	for _, v := range number {
		if v < '0' || v > '9' {
			return false
		}
	}

	// Apply Luhn algorithm
	var sum int
	ln := len(number)
	d := false

	// Iterate through string from right to left
	for i := ln - 1; i >= 0; i-- {

		n, err := strconv.Atoi(string(number[i]))
		if err != nil {
			log.Fatalln("Error converting string to integer:", err)
		}

		// Double n every second iteration, if sum is greater than 9, subtract 9
		if d {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		d = !d

		sum += n
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
