package consecutive_strings

import (
	 "testing"
	"log"
	"fmt"
)

func TestLongestConsec(t *testing.T) {

	cases := []struct{
		strings []string
		consec int
		expect string
	} {
		{[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta"},
		{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh"},
		{[]string{}, 3, ""},
		{[]string{"itvayloxrp","wkppqsztdkmvcuwvereiupccauycnjutlv","vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"},
	}

	for _, c := range cases {
		result := LongestConsec(c.strings, c.consec)
		if result != c.expect {
			log.Fatal(fmt.Sprintf("Expected %s, got %s", c.expect, result))
		}
	}
}

