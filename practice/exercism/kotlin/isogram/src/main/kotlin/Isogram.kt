object Isogram {
    fun isIsogram(str: String): Boolean {
        var charCount = mutableMapOf<Char, Int>()
        for (ch in str.toLowerCase()) {
            var c = charCount.getOrDefault(ch, 0)
            if (c > 0 && ch != ' ' && ch != '-') return false
            charCount.put(ch, c + 1)
        }
        return true
    }
}