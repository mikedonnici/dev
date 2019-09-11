// Package accumulate provides an Accumulate function.
package accumulate

// Accumulate returns a slice of converted values.
func Accumulate(input []string, converter func(string) string) []string {
	result := make([]string, len(input))
	for i, s := range input {
		result[i] = converter(s)
	}
	return result
}
