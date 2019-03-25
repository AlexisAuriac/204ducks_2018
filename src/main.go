package main

import (
	"os"
	"fmt"
)

func main() {
	a := parseArgv(os.Args[1:])
	// var a float64 = parseArgv(os.Args[1:])

	fmt.Println("a:", a)
}
