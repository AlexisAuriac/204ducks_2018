package main

import "math"

// MeanDucks computes the mean of the integral in seconds.
func MeanDucks(a float64, tmax float64) uint {
	var inc = 0.001
	var mean = 0.0
	// t: time in milliseconds
	var t = 0.0

	for t < tmax {
		mean += DuckReturnProb(a, t) * t
		t += inc
	}

	// accounting for the augmentation of values
	// threw the fragmentation of minutes
	mean /= 1 / inc
	// minutes to seconds
	mean *= 60
	return uint(math.Ceil(mean))
}
