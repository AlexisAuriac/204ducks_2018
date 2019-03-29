package main

import (
	"fmt"
	"math"
	"os"
)

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

// Takes a proportion of the ducks and returns the time in seconds
// before this proportion is back
func FindTimeProportionDucksBack(a float64, p float64) uint {
	var diff float64
	var lastDiff float64 = 1

	for i := uint(0); ; i++ {
		diff = math.Abs(p - ProbDensity(a, float64(i)/60))
		if diff > lastDiff {
			return i - 1
		}
		lastDiff = diff
	}
}

func main() {
	var a float64 = parseArgv(os.Args[1:])

	var t uint = FindTimeProportionDucksBack(a, 0.5)
	fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n",
		t/60, t%60)
	t = FindTimeProportionDucksBack(a, 0.99)
	fmt.Printf("Time after which 99%% of the ducks are back: %dm %02ds\n",
		t/60, t%60)

	fmt.Printf("Percentage of ducks back after 1 minute: %.1f%%\n",
		ProbDensity(a, 1)*100)
	fmt.Printf("Percentage of ducks back after 2 minute: %.1f%%\n",
		ProbDensity(a, 2)*100)
}
