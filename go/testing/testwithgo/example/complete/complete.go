// Package complete is an example on generating a full example with godoc
package complete

import "fmt"

func ShowExample(foo string) string {
	return fmt.Sprintf("A complete example of %s", foo)
}
