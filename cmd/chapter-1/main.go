package main

import (
	"bufio"
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

// Echoing line of texts
func echoLine() {
	scanner := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "bye" {
			break
		}

		counts[line] += 1
	}
	for line, count := range counts {
		fmt.Printf("%s: %d\n", line, count)
	}
}

// Finding duplicate lines
func dup2() {
	counts := make(map[string]int)
	args := os.Args
	// read from STDIN
	if len(args) == 1 {
		countLines(os.Stdin, counts)
	} else { // read from files
		for _, fileName := range args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("Error in opening %s\n", fileName)
				continue
			}
			countLines(file, counts)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counts[line] += 1
	}

	for line, count := range counts {
		fmt.Printf("%s: %d \n", line, count)
	}
}

func main() {
	//Exercise11()
	//fmt.Println(" ==========================")
	//Exercise12()
	//echoLine()
	dup2()
}
