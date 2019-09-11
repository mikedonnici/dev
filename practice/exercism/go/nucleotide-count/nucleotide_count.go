// Package dna provides a nucleotide count function
package dna

import (
	"errors"
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i, n := range d {
		if _, valid := h[n]; !valid {
			return h, errors.New(fmt.Sprintf("Invalid nucleotide '%v' at position %d", n, i+1))
		}
		h[n] += 1
	}

	return h, nil
}
