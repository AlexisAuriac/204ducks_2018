package main

import (
	"fmt"
	"os"
	"math"
)

func duckReturnProb(a float64, t uint) float64 {
	return (a*math.Exp(-float64(t)) + (4-3*a)*math.Exp(-2*float64(t)) + (2*a-4)*math.Exp(-4*float64(t)))
}

func primDuckReturnProb(a float64, t uint) float64 {
	return (-a*math.Exp(-float64(t)) - (4-3*a)/2*math.Exp(-2*float64(t)) - (2*a-4)/4*math.Exp(-4*float64(t)))
}

func probDensity(a float64, t uint) float64 {
	return (primDuckReturnProb(a, t) - primDuckReturnProb(a, 0)) * 100
}

func main() {
	var a float64 = parseArgv(os.Args[1:])

	fmt.Printf("Percentage of ducks back after 1 minute: %.1f\n", probDensity(a, 1))
	fmt.Printf("Percentage of ducks back after 2 minute: %.1f%%\n", probDensity(a, 2))
}
