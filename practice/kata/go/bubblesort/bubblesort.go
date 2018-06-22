package bubblesort

// BubbleSort runs through a slice of numbers once for each element in the slice
// and pushed the largest number to its correct position on each pass.
func BubbleSort(numbers []int) []int {

	l := len(numbers)

	for i := 0; i < l; i++ {
		for j := 0; j < (l-1-i); j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}
