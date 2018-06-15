// Package bob creates lackadaisical teenager responses :)
package bob

import (
	"strings"
)

// Hey determines an appropriate response based on the input string
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	switch {
	case allCaps(remark) && endsWithQuestionMark(remark):
		// shouting and question
		return "Calm down, I know what I'm doing!"

	case endsWithQuestionMark(remark):
		// question
		return "Sure."

	case allCaps(remark):
		// shouting
		return "Whoa, chill out!"

	case len(remark) == 0:
		// silent address
		return "Fine. Be that way!"

	default:
		return "Whatever."
	}
}

// allCaps returns true if s contains only uppercase characters and at least one alpha character
func allCaps(s string) bool {
	return (s == strings.ToUpper(s)) && containsAlpha(s)
}

// containsAlpha returns true is any alpha character is found in the string
func containsAlpha(s string) bool {
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

// endsWithQuestionMark returns true if the string ends with '?'
func endsWithQuestionMark(s string) bool {
	if len(s) > 0 {
		l := s[len(s)-1]
		if l == '?' {
			return true
		}
	}
	return false
}
