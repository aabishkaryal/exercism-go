package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides classifies the traingle type given the length of sides of triangle.
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

// fulfilsTriangleInequality checks if the given triangle fulfils the triangle inequality.
func fulfilsTriangleInequality(a, b, c float64) bool {
	return 2*max(a, b, c) <= a+b+c
}

// max returns the maximum value from given numbers
func max(nums ...float64) float64 {
	l := nums[0]
	for _, v := range nums[1:] {
		if v > l {
			l = v
		}
	}
	return l
}

// containsInf checks if nums contains a infinite value.
func containsInf(nums ...float64) bool {
	for _, v := range nums {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// containsZero checks if nums contains a zero.
func containsZero(nums ...float64) bool {
	for _, v := range nums {
		if v == 0 {
			return true
		}
	}
	return false
}

// containsNegative checks if nums contains a negative number.
func containsNegative(nums ...float64) bool {
	for _, v := range nums {
		if v < 0 {
			return true
		}
	}
	return false
}

// containsNan checks if nums contains a Nan.
func containsNan(nums ...float64) bool {
	for _, v := range nums {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}
