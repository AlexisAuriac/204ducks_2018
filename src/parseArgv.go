package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func usage() {
	fmt.Println("USAGE")
	fmt.Println("\t./204ducks\ta")
	fmt.Println()
	fmt.Println("DESCRIPTION")
	fmt.Println("\ta\t\tconstant")
}

func parseArgv(argv []string) float64 {
	if len(argv) != 1 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(84)
	} else if argv[0] == "-h" {
		usage()
		os.Exit(0)
	}

	r, _ := regexp.Compile("^\\d+(\\.\\d+)?$")

	if !r.MatchString(argv[0]) {
		fmt.Fprintln(os.Stderr, argv[0], ": Invalid constant")
		os.Exit(84)
	}

	a, _ := strconv.ParseFloat(argv[0], 64)

	if a < 0 || a > 2.5 {
		fmt.Fprintln(os.Stderr, argv[0], ": Invalid constant")
		os.Exit(84)
	}

	return a
}
