package main

import (
	"fmt"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	values := make([]bool, n*n)
	var repeated, missing int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			value := grid[i][j] - 1
			if values[value] {
				repeated = value + 1
			}
			values[value] = true
		}
	}
	for i := 0; i < n*n; i++ {
		if !values[i] {
			missing = i + 1
			break
		}
	}
	return []int{repeated, missing}
}

func main() {
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	result := findMissingAndRepeatedValues(grid)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[9 5]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
