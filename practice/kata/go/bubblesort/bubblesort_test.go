package bubblesort

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		input []int
		expect []int
	}{
		{input: []int{3,2,1}, expect: []int{1,2,3}},
		{input: []int{1,2,3}, expect: []int{1,2,3}},
		{input: []int{1,2,1}, expect: []int{1,1,2}},
		{input: []int{4,2,1}, expect: []int{1,2,4}},
		{input: []int{1,5,2,7,9,4}, expect: []int{1,2,4,5,7,9}},
	}

	for _, c := range cases {
		xi := BubbleSort(c.input)
		if !reflect.DeepEqual(xi, c.expect) {
			log.Fatal(fmt.Sprintf("input %v, expect %v", c.input, c.expect))
		}
	}
}
