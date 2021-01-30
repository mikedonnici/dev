package adder

func Add(a, b int) int {
	return a + b
}

func Sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
