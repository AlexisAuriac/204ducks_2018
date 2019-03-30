package main

import "math"

// Computes the proportion of ducks that will be back after t minutes,
// given the constant a
func DuckReturnProb(a float64, t float64) float64 {
	return (a*math.Exp(-t) +
		(4-3*a)*math.Exp(-2*t) +
		(2*a-4)*math.Exp(-4*t))
}

// Primitive of the formula defined above
func PrimDuckReturnProb(a float64, t float64) float64 {
	return (-a*math.Exp(-t) -
		(4-3*a)/2*math.Exp(-2*t) -
		(2*a-4)/4*math.Exp(-4*t))
}

// Rerturns the probability density of the integral (from 0 to t)
func ProbDensity(a float64, t float64) float64 {
	return PrimDuckReturnProb(a, t) - PrimDuckReturnProb(a, 0)
}
