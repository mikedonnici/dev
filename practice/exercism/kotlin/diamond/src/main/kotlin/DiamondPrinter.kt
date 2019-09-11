class DiamondPrinter {

  fun printToList(ch: Char): List<String> {

    if (ch == 'A') return listOf("A")

    val range = 'A'..ch
    val width = range.count() * 2 - 1
    val centreIndex = range.count() - 1
    var leftIndex = centreIndex
    var rightIndex = centreIndex
    var result = mutableListOf<String>()

    for (char in range) {

      var row = ""
      for (j in 0..width - 1) {
          if (j == leftIndex || j == rightIndex) {
              row += char
          } else {
              row += " "
          }
      }

      leftIndex--
      rightIndex++
      result.add(row)
    }

    return (result + result.reversed().subList(1, range.count()))
  }
}

