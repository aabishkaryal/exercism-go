package darts

import "math"

// Score calculates score for the given x, y cordinate in darts game.
// - If the dart lands outside the target, player earns no points (0 points).
// - If the dart lands in the outer circle of the target, player earns 1 point.
// - If the dart lands in the middle circle of the target, player earns 5 points.
// - If the dart lands in the inner circle of the target, player earns 10 points.
func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	if distance > 10 {
		return 0
	}
	if distance > 5 {
		return 1
	}
	if distance > 1 {
		return 5
	}
	return 10
}
