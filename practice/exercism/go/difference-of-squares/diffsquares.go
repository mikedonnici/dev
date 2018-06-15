package diffsquares

// SquareOfSums calculates the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	var sm int
	for i := 1; i <= n; i++ {
		sm += i
	}
	return sm * sm
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	var sm int
	for i := 1; i <= n; i++ {
		sq := i * i
		sm += sq
	}
	return sm
}

// Difference returns the difference between the sum of squares and the square of the sum, of the first n natural numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
