// Package triangle provides functions for triangles
package triangle

import "math"

// Kind distinguishes the type of triangle
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines the type of triangle based on the dimensions a, b and c.
func KindFromSides(a, b, c float64) Kind {

	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c && c == a {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

// isTriangle checks that the dimension can create a triangle
func isTriangle(a, b, c float64) bool {
	return dimensionsOK(a, b, c) && sumDimensionsOK(a, b, c)
}

func dimensionsOK(a, b, c float64) bool {
	return a > 0 && !math.IsInf(a, 1) &&
		b > 0 && !math.IsInf(b, 1) &&
		c > 0 && !math.IsInf(c, 1)
}

// sum of the two smaller sides must exceed third side
func sumDimensionsOK(a, b, c float64) bool {
	if a > b && a > c {
		return b+c >= a
	}
	if b > c && b > a {
		return a+c >= b
	}
	return a+b >= c
}
