package anagram

import "strings"

// Detect returns one or more anagrams of the subject from a list of candidates.
func Detect(subject string, candidates []string) []string {
	var anagrams []string
	if subject == "" || len(candidates) == 0 {
		return anagrams
	}
	for _, c := range candidates {
		if isAnagram(subject, c) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func isAnagram(a, b string) bool {

	if len(a) != len(b) {
		return false
	}
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	aMap := make(map[rune]bool)
	for _, c := range a {
		aMap[c] = true
	}

	bMap := make(map[rune]bool)
	for _, c := range b {
		bMap[c] = true
	}

	for _, c := range b {

		for _, c2 := range b {
			if c1 == c2 {
				continue out
			}
		}
		return false
	}
	return true
}