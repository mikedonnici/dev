package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//

func DecToBase(dec, base int) string {

	var result string

	// divide the number by the base - the remainder is the number that belongs in the current position
	result = numString(dec % base) + result


	// quotient less than 1 we are done - no fractions here
	quotient := dec / base
	if quotient < 1 {
		return result
	}

	// recursive call to prepend the next number to the final result
	return DecToBase(quotient, base) + result
}

// numString returns the string representation of a single digit number up to base 16
func numString(n int) string {
	return fmt.Sprintf("%X", n)
}
