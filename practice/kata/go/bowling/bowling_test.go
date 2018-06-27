package bowling

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {

	cases := []struct {
		bowls []int
		score int
	}{
		{
			bowls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			score: 0,
		},
		{
			bowls: []int{0, 1, 3, 5, 1, 4, 7, 8, 9, 1},
			score: 39,
		},
		{
			bowls: []int{1, 9, 3},
			score: 16,
		},
		{
			bowls: []int{10, 9},
			score: 28,
		},
		{
			bowls: []int{1, 9, 3, 6, 7},
			score: 29,
		},
		{
			bowls: []int{7, 3, 3, 7, 7},
			score: 37,
		},
		{
			bowls: []int{1,9,1,9,1,9,1,9,1,9,1,9,1,9,1,9,1,9,1,9,1},
			score: 110,
		},
		{
			bowls: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			score: 300,
		},
	}

	for _, c := range cases {
		s := Score(c.bowls)
		if s != c.score {
			t.Fatal(fmt.Sprintf("Expected score %v, got %v", c.score, s))
		}
	}
}
