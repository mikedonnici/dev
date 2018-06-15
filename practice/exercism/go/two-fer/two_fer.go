// Package twofer does the 'twofer' thing
package twofer

import "fmt"

// ShareWith returns a string that includes the name passed in, or a generic string
func ShareWith(n string) string {
	name := "you"
	if len(n) > 0 {
		name = n
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
