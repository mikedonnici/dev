package fibonacci

func FibonacciRecursive(n int) int {

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return FibonacciRecursive(n - 2) + FibonacciRecursive(n - 1)
}
