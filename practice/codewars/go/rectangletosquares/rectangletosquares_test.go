package rectangletosquares

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {

	cases := []struct {
		l    int
		w    int
		want []int
	}{
		{1, 1, nil},
		{5, 5, nil},
		{2, 1, []int{1, 1}},
		{5, 3, []int{3, 2, 1, 1}},
	}

	for _, c := range cases {
		got := SquaresInRect(c.l, c.w)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SquaresInRect(%v, %v) = %v, want %v", c.l, c.w, got, c.want)
		}
	}
}
