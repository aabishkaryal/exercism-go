// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if containsZero(a, b, c) || containsNegative(a, b, c) || containsInf(a, b, c) || containsNan(a, b, c) || !fulfilsTriangleInequality(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}

func fulfilsTriangleInequality(a, b, c float64) bool {
	return 2*max(a, b, c) <= a+b+c
}

func max(nums ...float64) float64 {
	l := nums[0]
	for _, v := range nums[1:] {
		if v > l {
			l = v
		}
	}
	return l
}

func containsInf(nums ...float64) bool {
	for _, v := range nums {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

func containsZero(nums ...float64) bool {
	for _, v := range nums {
		if v == 0 {
			return true
		}
	}
	return false
}

func containsNegative(nums ...float64) bool {
	for _, v := range nums {
		if v < 0 {
			return true
		}
	}
	return false
}

func containsNan(nums ...float64) bool {
	for _, v := range nums {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}
