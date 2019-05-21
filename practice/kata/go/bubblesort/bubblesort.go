package bubblesort

// BubbleSort sorts a slice of int.
// Note this version is LESS efficient than the previous as it
// iterates through the entire slice for each comparison
// more efficient version: https://play.golang.org/p/_pmWoHabfgq
func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[j] > numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}
