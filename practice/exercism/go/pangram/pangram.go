package pangram

import "strings"

// IsPangran returns true if the string is a pangram.
func IsPangram(str string) bool {
	chars := make(map[rune]bool)
	str = strings.ToLower(str)
	for _, ch := range str {
		if ch >= 97 && ch <= 122 {
			chars[ch] = true
		}
	}
	return len(chars) == 26
}
