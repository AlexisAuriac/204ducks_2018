package main

import (
	"fmt"
	"os"
)

func main() {
	var a = ParseArgv(os.Args[1:])

	var t50 = FindTimeDucksBack(a, 0.5)
	var tMax = FindTimeDucksBack(a, 0.99)
	var mean = MeanDucks(a, float64(tMax))
	var stdDev = StdDevDucks(a, float64(tMax), mean)

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
