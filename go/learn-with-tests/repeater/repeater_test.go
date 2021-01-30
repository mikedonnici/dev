package repeater

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("Repeat() = %s, want %s", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 5)
	fmt.Print(s)
	// Output: aaaaa
}
