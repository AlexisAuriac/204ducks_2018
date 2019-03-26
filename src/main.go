package main

import (
	"os"
	"fmt"
)

func main() {
	var a float64 = parseArgv(os.Args[1:])

	fmt.Printf("Percentage of ducks back after 1 minute: %.1f\n", duckPercentage(a, 1))
	fmt.Printf("Percentage of ducks back after 2 minute: %.1f\n", duckPercentage(a, 2))
}
