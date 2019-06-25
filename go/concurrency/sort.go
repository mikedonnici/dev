package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	nums, err := inputNumbers()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Original list:", nums)

	ch := make(chan int)

	// get sub slices
	xxi := subLists(nums, 4)
	fmt.Println("Sub lists:", xxi)

	for i := range xxi {
		go chanSort(xxi[i], ch, i+1)
	}

	var merged []int
	for range nums {
		n := <- ch
		merged = append(merged, n)
	}	
	fmt.Println("Merged:", merged)

	fmt.Println("Merged and sorted:", bubbleSort(merged))

}

// inputNumbers inputs the numbers from the users and returns a slice of int.
func inputNumbers() ([]int, error) {

	fmt.Print("Enter a list of integers (seperated by spaces): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()

	return numbersFromLine(l)
}

// numbersFromLine splits the string argument by whitespace, converts each
// element to an integer and returns the resulting slice int. It will return an
// error if any of the strings cannot be converted to an integer.
func numbersFromLine(s string) ([]int, error) {

	var result []int
	xs := strings.Fields(s)
	for _, s := range xs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return result, err
		}
		result = append(result, n)
	}

	return result, nil
}

// seperate splits a slice of integers into a number of sub slices
func subLists(nums []int, parts int) [][]int {

	// empty data structure with specified sub slices
	var result [][]int
	for i := 0; i < parts; i++ {
		result = append(result, []int{})
	}

	// fill it up
	c := 0
	for i := 0; i < len(nums); i++ {
		if c == parts {
			c = 0
		}
		result[c] = append(result[c], nums[i])
		c++
	}

	return result
}

// chanSort sorts a slice of integers and sends them 
// over the channel in ascending order 
func chanSort(nums []int, ch chan int, routineNumber int) {
	fmt.Print("goroutine", routineNumber, "sort list", nums, "->")
	nums = bubbleSort(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		ch <- nums[i]
	}
}

// bubblesort returns a slice integers in ascending order
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// insert inserts a value at an appropriate place according to its value
func insert(n int, nums []int) []int {

	// n is first element
	if len(nums) == 0 {
		return []int{n}
	}

	// n lowest so first
	if n < nums[0] {
		return append([]int{n}, nums...)
	}

	// n somewhere in the middle
	for i := range nums {
		if n >= nums[i] {
			var result []int
			head := nums[:i]
			tail := nums[i:]
			result = append(result, head...)
			result = append(result, n)
			result = append(result, tail...)
			return result
		}
	} 

	// n is last
	return append(nums, n) 
}