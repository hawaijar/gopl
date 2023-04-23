package main

import (
	"fmt"
	"os"
	"strings"
)

// Exercise11 (1.1) - Modify the echo program to also print os.Args[0], the name of the command that invoked it.
func Exercise11() {
	programName := os.Args[0]
	programArgs := strings.Join(os.Args[1:], " ")
	fmt.Println("Program name:", programName)
	fmt.Println("Program arguments:", programArgs)
}

// Exercise12 (1.2) - Modify the echo program to print the index and value of each of its arguments, one per line.
func Exercise12() {
	programName := os.Args[0]
	fmt.Println("Program name:", programName)
	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, arg)
	}
}

// Exercise13 (1.3) - Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
func Exercise13() {

}

func main() {
	Exercise11()
	fmt.Println(" ==========================")
	Exercise12()
}
