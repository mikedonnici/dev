package main

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numLists ...[]int) []int {
	var sums []int
	for _, numList := range numLists {
		sums = append(sums, Sum(numList))
	}
	return sums
}

func SumAllTails(numLists ...[]int) []int {
	for i := 0; i < len(numLists); i++ {
		if len(numLists[i]) != 0 {
			numLists[i] = numLists[i][1:]
		}
	}
	return SumAll(numLists...)
}
