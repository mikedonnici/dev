package greet

import (
	"fmt"
	"testing"
)

func ExampleHello() {
	greeting := Hello("Mike")
	fmt.Println(greeting)
	// Output: Hello, Mike
}

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"Mike"}, "Hello, Mike"},
		{"test2", args{"Christie"}, "Hello, Christie"},
		{"test3", args{"Maia"}, "Hello, Maia"},
		{"test4", args{"Leo"}, "Hello, Leo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
