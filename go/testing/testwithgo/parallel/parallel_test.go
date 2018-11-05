package parallel

import (
	"testing"
)

func TestFoo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		{"testFoo", "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Log("TestFoo() complete")
}

func TestBar(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		{"testBar", "Foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bar(); got != tt.want {
				t.Errorf("Bar() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Log("TestBar() complete")
}

func TestNada(t *testing.T) {
	t.Log("TestNada() complete")
}
