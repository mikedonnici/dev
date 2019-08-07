// Package leap provides functionality for leap years
package leap

// IsLeapYear determines if a given year is a leap year
func IsLeapYear(year int) bool {

	div400 := year % 400 == 0
	div100 := year % 100 == 0
	div4 := year % 4 == 0

	return  (div4 && div400) || (div4 && !div100)
}
