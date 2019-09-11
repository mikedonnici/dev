// Package protein provides translation of RNA to protein sequences.
package protein

import (
	"errors"
)

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("Stop error")
var ErrInvalidBase = errors.New("Invalid base")

// FromRNA translates RNA to polypeptide sequences
func FromRNA(rna string) ([]string, error) {

	var proteinSequence []string

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteinSequence, nil
		}
		if err != nil {
			return proteinSequence, err
		}
		proteinSequence = append(proteinSequence, protein)
	}

	return proteinSequence, nil
}

// FromCodon returns the protein that corresponds to the codon.
func FromCodon(codon string) (string, error) {

	protein, ok := codonToProtein[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
