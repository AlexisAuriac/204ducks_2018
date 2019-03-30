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

// Computes the mean of the integral in seconds.
func MeanDucks(a float64, tmax float64) uint {
	var inc float64 = 0.001
	var mean float64 = 0
	// t: time in milliseconds
	var t float64 = 0

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

func main() {
	var a float64 = parseArgv(os.Args[1:])

	var t50 uint = FindTimeProportionDucksBack(a, 0.5)
	var tMax uint = FindTimeProportionDucksBack(a, 0.99)
	var mean uint = MeanDucks(a, float64(tMax))
	var stdDev float64 = StdDevDucks(a, float64(tMax), mean)

	fmt.Printf("Average return time: %dm %02ds\n", mean/60, mean%60)
	fmt.Printf("Standard deviation: %.3f\n", stdDev)

	fmt.Printf("Time after which 50%% of the ducks are back: %dm %02ds\n",
		t50/60, t50%60)
	fmt.Printf("Time after which 99%% of the ducks are back: %dm %02ds\n",
		tMax/60, tMax%60)

	fmt.Printf("Percentage of ducks back after 1 minute: %.1f%%\n",
		ProbDensity(a, 1)*100)
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f%%\n",
		ProbDensity(a, 2)*100)
}
