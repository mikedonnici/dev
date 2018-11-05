package parallelgotcha

import (
	"fmt"
	"testing"
)

func TestSquare_nofail(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{1, 1},
		{2, 5}, // should fail
		{3, 9},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel() // this causes subtest func call onto a queue and subtest exits
			t.Logf("arg=%d, want=%d", tc.arg, tc.want)
			if Square(tc.arg) != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}

func TestSquare_fail(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{1, 1},
		{2, 5}, // should fail
		{3, 9},
	}
	for _, tc := range testCases {
		tc := tc // create shadow (local copy) of tc for closure
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel() // this causes subtest func call onto a queue and subtest exits
			t.Logf("arg=%d, want=%d", tc.arg, tc.want)
			if Square(tc.arg) != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}
