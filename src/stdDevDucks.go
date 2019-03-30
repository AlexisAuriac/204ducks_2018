package main

import "math"

// Computes the standard deviation of the integral in minutes.
func StdDevDucks(a float64, tmax float64, mean uint) float64 {
	var inc float64 = 0.001
	var variance float64 = 0
	var t float64 = 0
	var mean2 float64 = float64(mean) / 60

	for t < tmax {
		variance += DuckReturnProb(a, t) * math.Pow(t-mean2, 2)
		t += inc
	}

	// accounting for the augmentation of values
	// threw the fragmentation of minutes
	variance /= 1 / inc
	return math.Sqrt(variance)
}
