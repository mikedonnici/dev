// Package strand provides functions for working with nucleotides.
package strand

// ToRNA returns the transcribed RNA sequence for the provided DNA sequence.
func ToRNA(dna string) string {

	var rna string

	t := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	for _, n := range dna {
		rna += t[string(n)]
	}

	return rna
}
