// Package raindrops ...
package raindrops

import "strconv"

// Convert converts a number to a certain string, depending on factors.
func Convert(n int) string {
	s := ""

	if n%3 == 0 {
		s += "Pling"
	}

	if n%5 == 0 {
		s += "Plang"
	}

	if n%7 == 0 {
		s += "Plong"
	}

	if len(s) > 0 {
		return s
	}

	return strconv.Itoa(n)
}
