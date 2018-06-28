package bubblesort

import (
	"testing"
	"reflect"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct{
		input []int
		expect []int
	}{
		{
			input: []int{2,1},
			expect: []int{1,2},
		},
		{
			input: []int{9,8,7,6,5,4,3,2,1},
			expect: []int{1,2,3,4,5,6,7,8,9},
		},
		{
			input: []int{1,1,1,1,3,2,1,1,1,1},
			expect: []int{1,1,1,1,1,1,1,1,2,3},
		},
	}

	for _, c := range cases {
		sorted := BubbleSort(c.input)
		if !reflect.DeepEqual(sorted, c.expect) {
			t.Fatalf(fmt.Sprintf("For input %v, expected %v, got %v", c.input, c.expect, sorted))
		}
	}
}
