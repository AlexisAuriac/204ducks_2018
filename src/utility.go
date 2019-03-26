package main

import "math"

func duckPercentage(a float64, t uint) (p float64) {
	return (a * math.Exp(-float64(t)) + (4 - 3 * a) * math.Exp(-2 * float64(t)) + (2 * a - 4) * math.Exp(-4 * float64(t))) * 100
}
