package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const FAILURE int = 84
const SUCCESS int = 0

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
		os.Exit(FAILURE)
	} else if argv[0] == "-h" || argv[0] == "--help" {
		usage()
		os.Exit(SUCCESS)
	}

	r, _ := regexp.Compile("^\\d+(\\.\\d+)?$")

	if !r.MatchString(argv[0]) {
		fmt.Fprint(os.Stderr, argv[0], ": Invalid constant\n")
		os.Exit(FAILURE)
	}

	a, _ := strconv.ParseFloat(argv[0], 64)

	if a < 0 || a > 2.5 {
		fmt.Fprint(os.Stderr, argv[0], ": Invalid constant\n")
		os.Exit(FAILURE)
	}

	return a
}
