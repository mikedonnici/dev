// Package romannumerals converts arabic numbers to roman numerals
package romannumerals

import (
	"errors"
	"strconv"
)

// ToRomanNumeral converts an arabic numver to a roman numeral.
func ToRomanNumeral(number int) (string, error) {

	var result string

	if number < 1 {
		return result, errors.New("Cannot convert number less than 1")
	}
	if number > 3000 {
		return result, errors.New("Cannot convert number greater that 3000")
	}

	units := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands := []string{"M", "MM", "MMM"}
	roman := [][]string{units, tens, hundreds, thousands}

	numstr := strconv.Itoa(number)
	set := 0 // units, tens, hundreds, thousands
	for i := len(numstr) - 1; i >= 0; i-- {
		index, err := strconv.Atoi(string(numstr[i]))
		if err != nil {
			return result, err
		}
		if index > 0 {
			result = roman[set][index-1] + result
		}
		set++
	}

	return result, nil
}
