// Package hamming calculates the Hamming distance between two strands of DNA
package hamming

import "github.com/pkg/errors"

// Distance receives two strings (a,b) representing DNA nucleotide sequences, and returns the number of
// nucleotides that differ between the two strands of DNA (Hamming distance).
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Strands must be equal length")
	}

	hd := 0
	for i := range a {
		if a[i] != b[i] {
			hd++
		}
	}

	return hd, nil
}
