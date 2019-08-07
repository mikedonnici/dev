// Package proverb provides the Proverb function
package proverb

import "fmt"

const line = "For want of a %s the %s was lost."
const end = "And all for the want of a %s."

// Proverb generates a proverb from the input
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return nil
	}

	var prov []string
	for i := 0; i < len(rhyme)-1; i++ {
		prov = append(prov, fmt.Sprintf(line, rhyme[i], rhyme[i+1]))
	}

	return append(prov, fmt.Sprintf(end, rhyme[0]))
}
