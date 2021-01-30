package hello

import "testing"

func TestHello(t *testing.T) {
	const f = "Hello()"
	assertEqual := func(t *testing.T, funcName string, got string, want string) {
		t.Helper() // ensures failures report the correct line number!
		if got != want {
			t.Errorf("%s = %q, want %q", funcName, got, want)
		}
	}

	t.Run("Say hello with string", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike!"
		assertEqual(t, f, got, want)
	})

	t.Run("Say hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertEqual(t, f, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("Sibelle", "ES")
		want := "Hola, Sibelle!"
		assertEqual(t, f, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("Marcelle", "FR")
		want := "Bonjour, Marcelle!"
		assertEqual(t, f, got, want)
	})

	t.Run("Say hello in Italian", func(t *testing.T) {
		got := Hello("Joseph", "ITALIAN")
		want := "Bonjourno, Joseph!"
		assertEqual(t, f, got, want)
	})
}
