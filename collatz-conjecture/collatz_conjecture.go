package collatzconjecture

import "errors"

// Function CollatzConjecture returns the number of steps needed for n to reach 1 in a colltaz series.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("provided number should be non-zero positive number")
	}
	step := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		step++
	}
	return step, nil
}
