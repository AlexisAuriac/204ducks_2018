package main

import "math"

// StdDevDucks computes the standard deviation of the integral in minutes.
func StdDevDucks(a float64, tmax float64, mean uint) float64 {
	var inc = 0.001
	var variance = 0.0
	var t = 0.0
	var mean2 = float64(mean) / 60

	for t < tmax {
		variance += DuckReturnProb(a, t) * math.Pow(t-mean2, 2)
		t += inc
	}

	// accounting for the augmentation of values
	// threw the fragmentation of minutes
	variance *= inc
	return math.Sqrt(variance)
}
