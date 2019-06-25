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

	fmt.Println(nums)

	xxi := subLists(nums, 4)
	fmt.Println(xxi)

}

// inputNumbers inputs the numbers from the users and returns a slice of int.
func inputNumbers() ([]int, error) {

	fmt.Print("Enter a list of integers (seperated by spaces): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()

	// numLists = subLists(nums, 4)

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
