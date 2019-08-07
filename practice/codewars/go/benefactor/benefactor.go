package benefactor

import (
	"math"
)

func NewAvg(arr []float64, navg float64) int64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	n := (navg * float64((len(arr) + 1))) - sum
	if n < 0 {
		return -1
	}
	return int64(math.Ceil(n))
}
