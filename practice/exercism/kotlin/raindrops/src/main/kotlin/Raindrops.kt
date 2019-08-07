object Raindrops {
    fun convert(num: Int): String {
        var str = if (num % 3 == 0) "Pling" else ""
        str += if (num % 5 == 0) "Plang" else ""
        str += if (num % 7 == 0) "Plong" else ""
        return if (str.length > 0) str else num.toString()
    }
}