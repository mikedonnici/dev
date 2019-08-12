object Hamming {
    fun compute(dna1: String, dna2: String): Int {
        var hdist = 0

        if (dna1.length != dna2.length) {
            throw IllegalArgumentException("left and right strands must be of equal length.")
        }

        for (i in dna1.indices) {
            if (dna1[i] != dna2[i]) hdist += 1
        }

        return hdist
    }
}