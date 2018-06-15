// Package isogram determines if strings are isograms.
package isogram

import "strings"

// IsIsogram determines if a string is an isogram.
func IsIsogram(s string) bool {

	s = strings.ToLower(s)

	for r := 'a'; r <= 'z'; r++ {
		if strings.Count(s, string(r)) > 1 {
			return false
		}
	}

	return true
}
