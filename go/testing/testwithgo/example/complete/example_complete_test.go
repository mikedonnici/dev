package complete_test

import (
	"fmt"
	"github.com/mikedonnici/dev/go/testing/testwithgo/example/complete"
)

const notUsed = "required for the full example to be generated in code"

// Package-level example
func Example() {
	fmt.Print(complete.ShowExample("Something"))
	// Output: A complete example of Something
}
