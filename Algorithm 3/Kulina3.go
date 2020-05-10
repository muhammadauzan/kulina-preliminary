package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Declaration
	var textnumber string

	// Scan console one line and formatting
	fmt.Scanln(&textnumber)
	textnumber = strings.Replace(textnumber, ".", "", -1)

	// Determine input properties
	length := len(textnumber)
	number := make([]int, length)

	// Calculate decimal digit
	for i := 0; i < length; i++ {
		number[i] = int(textnumber[i]-48) * int(math.Pow10(length-i))
		fmt.Println(number[i])
	}
}
