class Anagram(val subject: String) {

  fun match(list: List<String>): Set<String> {
    val anagrams = mutableSetOf<String>()
    list.forEach { if (isAnagram(subject, it)) anagrams.add(it) }
    return anagrams
  }

  fun isAnagram(subject: String, candidate: String): Boolean {
    return (
      subject.toLowerCase().all { it -> candidate.toLowerCase().contains(it) } &&
      subject.length == candidate.length
    )
  }
}
