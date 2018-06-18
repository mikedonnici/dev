package letter

type FreqMap map[rune]int

// Frequency counts the occurrences of each letter in a string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts occurrences of letters in a number of strings, concurrently.
func ConcurrentFrequency(xs []string) FreqMap {
	ch := make(chan FreqMap, len(xs))
	fm := FreqMap{}
	for _, s := range xs {
		go freq(s, ch)
	}
	for range xs {
		for r, c := range <-ch {
			fm[r] += c
		}
	}
	return fm
}

func freq(s string, ch chan FreqMap) {
	ch <- Frequency(s)
}
