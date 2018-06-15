// Package acronym handles acronyms
package acronym

import "strings"

// Abbreviate returns an acronym based on the string argument
func Abbreviate(s string) string {

	a := ""

	// replace hyphens with space and separate words
	s = strings.Replace(s, "-", " ", -1)
	xs := strings.Split(s, " ")

	for _, v := range xs {
		v = strings.TrimSpace(v)
		a = a + string(v[0])
	}

	return strings.ToUpper(a)
}
