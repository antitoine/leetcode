package main

import (
	"fmt"
)

func coloredCells(n int) int64 {
	return int64(2*n*(n-1) + 1)
}

func main() {
	n := 2
	result := coloredCells(n)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
