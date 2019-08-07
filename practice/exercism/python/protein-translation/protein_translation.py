class ProteinTranslation:
    __codonToProtein = {
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

    def __init__(self, rna):
        self.__rna = rna

    def condonSequence(self):
        return [self.__rna[i:i + 3] for i in range(0, len(self.__rna), 3)]

    def aminoSequence(self):
        s = []
        codons = self.condonSequence()
        for c in codons:
            if c == "STOP":
                break
            s.append(self.__codonToProtein[c])
        return s


def proteins(strand):
    return ProteinTranslation("AUGUUUUCU").aminoSequence()
