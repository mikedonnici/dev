package bubblesort

import (
	"testing"
	"fmt"
	"reflect"
)

func TestBubbleSort(t *testing.T) {

	cases := []struct {
		input  []int
		expect []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 4, 2, 5, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		result := BubbleSort(c.input)
		if !reflect.DeepEqual(result, c.expect) {
			t.Fatalf(fmt.Sprintf("For input %v expected %v, got %v", c.input, c.expect, result))
		}
	}
}
