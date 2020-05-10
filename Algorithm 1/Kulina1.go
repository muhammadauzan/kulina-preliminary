package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func matching(arr []int, num int) int {
	// Local variable declaration
	var value int
	var count int

	// Sorting sock array
	sort.Ints(arr)

	// Matching count algorithm
	for i := 0; i < num; i++ {
		if value != arr[i] {
			value = arr[i]
		} else {
			count++
			value = 0
		}
	}
	return count
}

func main() {
	// Declaration
	var socknumber int
	reader := bufio.NewReader(os.Stdin)

	// Accept number of sock input to make sock array
	fmt.Scanln(&socknumber)
	sock := make([]int, socknumber)

	// Accept sock number/name untill end line
	sockinput, _ := reader.ReadString('\n')

	// Remove end line
	sockinput = strings.Replace(sockinput, "\r\n", "", -1)

	// Split string to string array
	sockarray := strings.Split(sockinput, " ")

	// Convert string array to int array
	for i := 0; i < socknumber; i++ {
		num, err := strconv.Atoi(sockarray[i])
		if err == nil {
			sock[i] = num
		}
	}

	// Matching function
	matched := matching(sock, socknumber)
	fmt.Println(matched)
}
