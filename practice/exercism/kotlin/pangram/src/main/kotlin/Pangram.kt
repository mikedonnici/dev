object Pangram {
    fun isPangram(str: String) = "abcdefghijklmnopqrstuvwxyz".all { it ->
        str.toLowerCase().contains(it)
    }
}