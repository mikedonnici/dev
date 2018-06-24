package fizzbuzz

import (
	"fmt"
	"log"
	"testing"
)

func TestDivisibleByThree(t *testing.T) {
	cases := []struct {
		n int
		e bool
	}{
		{n: 3, e: true},
		{n: 1, e: false},
		{n: 9, e: true},
		{n: 0, e: true},
		{n: -12, e: true},
	}

	for _, c := range cases {
		if DivisibleByThree(c.n) != c.e {
			log.Fatal(fmt.Sprintf("%v should be divisible by 3, got %v", c.n, c.e))
		}
	}
}

func TestDivisibleByFive(t *testing.T) {
	cases := []struct {
		n int
		e bool
	}{
		{n: 0, e: true},
		{n: 4, e: false},
		{n: 15, e: true},
		{n: -5, e: true},
		{n: -6, e: false},
	}

	for _, c := range cases {
		if DivisibleByFive(c.n) != c.e {
			log.Fatal(fmt.Sprintf("%v should be divisible by 5, got %v", c.n, c.e))
		}
	}
}

func TestDivisibleByThreeAndFive(t *testing.T) {
	cases := []struct {
		n int
		e bool
	}{
		{n: 0, e: true},
		{n: 15, e: true},
		{n: 14, e: false},
		{n: -15, e: true},
	}

	for _, c := range cases {
		if DivisibleByThreeAndFive(c.n) != c.e {
			log.Fatal(fmt.Sprintf("%v should be divisible by 3 and 5, got %v", c.n, c.e))
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n      int
		expect string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{14, "14"},
		{18, "Fizz"},
		{21, "Fizz"},
		{25, "Buzz"},
		{30, "FizzBuzz"},
	}

	for _, c := range cases {
		s := FizzBuzz(c.n)
		if s != c.expect {
			t.Fatal(fmt.Sprintf("Expected %s, got %s", c.expect, s))
		}
	}
}
