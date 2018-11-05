package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{1,2,3})
	if sum != 6 {
		t.Errorf("wanted 6, got %d", sum)
	}
}
