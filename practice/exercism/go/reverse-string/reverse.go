// Package reverse reverses strings.
package reverse

import "fmt"

// String reverses a string.
func String(s string) string {
	rs := ""
	for _, r := range s {
		rs = fmt.Sprintf("%c", r) + rs
	}

	return rs
}
