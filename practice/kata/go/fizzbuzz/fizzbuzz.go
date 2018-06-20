package fizzbuzz

func DivisibleByThree(n int) bool {
	return n%3 == 0
}

func DivisibleByFive(n int) bool {
	return n%5 == 0
}

func DivisibleByThreeAndFive(n int) bool {
	return DivisibleByThree(n) && DivisibleByFive(n)
}