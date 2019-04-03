package main

import "math"

// FindTimeDucksBack takes a proportion of the ducks and returns the time in seconds
// before this proportion is back
func FindTimeDucksBack(a float64, p float64) uint {
	var diff float64
	var lastDiff float64 = 1

	for i := uint(0); ; i++ {
		diff = math.Abs(p - ProbDensity(a, float64(i)/60))
		if diff > lastDiff {
			return i - 1
		}
		lastDiff = diff
	}
	return 0
}
