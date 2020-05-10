package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countingValleys(arr []string, num int) int {
	// Local Declaration
	var value int
	var previous int
	var count int

	// Valley Counting Algorithm
	for i := 0; i < num; i++ {
		// Calculate Altitude based on step
		previous = value
		if arr[i] == "D" {
			value--
		} else {
			value++
		}

		// Determine if already past valley and count valley
		if value == 0 && previous < 0 {
			count++
		}
	}
	return count
}

func main() {
	// Declaration
	var stepsnumber int
	reader := bufio.NewReader(os.Stdin)

	// Accept number of steps input to make steps array
	fmt.Scanln(&stepsnumber)
	steps := make([]string, stepsnumber)

	// Accept steps U/D untill end line
	stepsinput, _ := reader.ReadString('\n')

	// Remove end line
	stepsinput = strings.Replace(stepsinput, "\r\n", "", -1)

	// Split string to string array
	steps = strings.Split(stepsinput, " ")

	valleys := countingValleys(steps, stepsnumber)
	fmt.Println(valleys)
}
