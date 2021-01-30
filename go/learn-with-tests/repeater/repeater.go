package repeater

func Repeat(s string, n int) string {
	var r string
	for i := 0; i < n; i++ {
		r += s
	}
	return r
}
