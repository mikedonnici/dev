package consecutive_strings

import "strings"

// LongestConsec returns the longest string comprised of k consecutive strings from the input slice
func LongestConsec(xs []string, k int) string {
	var cs string
	for i := 0; i <= len(xs) - k; i++ {
		s := strings.Join(xs[i:i+k], "")
		if len(s) > len(cs) {
			cs = s
		}
	}
	return cs
}
