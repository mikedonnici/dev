package module01

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	charToDecimal := map[string]float64 {
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
		"9": 9, "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15,
	}
	var result float64
	maxIndex := len(value)-1
	for i := maxIndex; i >= 0; i-- {
		pos := maxIndex - i
		result += charToDecimal[string(value[i])] * (math.Pow(float64(base), float64(pos)))
	}
	return int(result)
}
