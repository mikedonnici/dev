package benefactor

import (
	"testing"
)

func TestNewAvg(t *testing.T) {

	cases := []struct {
		arg1 []float64
		arg2 float64
		want int64
	}{
		{[]float64{1, 2, 3}, 7, 22},
		{[]float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 16.0}, 90, 628},
		{[]float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}, 92, 645},
	}

	for _, c := range cases {
		got := NewAvg(c.arg1, c.arg2)
		if got != c.want {
			t.Errorf("NewAvg() = %d, want %d", got, c.want)
		}
	}
}
