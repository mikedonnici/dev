// Package etl transforms legacy scrabble data
package etl

import "strings"

// Transform converts legacy scrabble score data to new format.
func Transform(input map[int][]string) map[string]int {

	res := make(map[string]int)

	for score, letters := range input {
		for _, l := range letters {
			res[strings.ToLower(l)] = score
		}
	}

	return res
}
