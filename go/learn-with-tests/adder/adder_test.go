package adder

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 1
	got := Add(a, b)
	want := 2
	if got != want {
		t.Errorf("Add(%d, %d) = %d, want %d", a, b, got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4
}

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3}
	got := Sum(nums)
	want := 6
	if got != want {
		t.Errorf("Sum() = %d, want %d", got, want)
	}
}
