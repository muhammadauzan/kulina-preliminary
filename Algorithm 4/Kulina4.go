package main

import "fmt"

func main() {
	// Declaration (switches number, and trips)
	num := 100
	trips := 100
	switches := make([]int, num)
	var count int

	// Simulate switch on/of for given trips and switches number
	for i := 0; i < trips; i++ {
		for j := 0 + i; j < num; j += (i + 1) {
			if switches[j] == 1 {
				switches[j] = 0
			} else {
				switches[j] = 1
			}
		}
	}

	// Calculate the number of turned on switches at the last trip
	for i := 0; i < len(switches); i++ {
		if switches[i] == 1 {
			count++
		}
	}

	// Display number of turned on switches. (For 100 switches and 100 trip, 10 still lamp turned on at the last trip)
	fmt.Println(count)
}
