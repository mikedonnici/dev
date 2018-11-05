package greet

import (
	"fmt"
	"strings"
)

// Hello returns a greeting string
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", strings.ToLower(name))
}
