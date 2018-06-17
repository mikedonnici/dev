package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(xs []string) {

	var ch chan FreqMap

	for _, s := range xs {
		go freq(s, ch)
	}
}

func freq(s string, ch chan FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
