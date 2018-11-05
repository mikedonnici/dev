package math

// Sum returns the sum of a slice of integers
func Sum(numbers []int) int {
	var sum int
	for n := range numbers {
		sum += n
	}
	return sum
}
