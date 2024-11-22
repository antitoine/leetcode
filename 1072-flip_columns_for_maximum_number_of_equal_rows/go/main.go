package main

import (
	"fmt"
)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patterns := make(map[string]int)
	for _, row := range matrix {
		var pattern, flippedPattern string
		for _, cell := range row {
			if cell == 1 {
				pattern += "1"
				flippedPattern += "0"
			} else {
				pattern += "0"
				flippedPattern += "1"
			}
		}
		if pattern < flippedPattern {
			patterns[pattern]++
		} else {
			patterns[flippedPattern]++
		}
	}

	maxRows := 0
	for _, count := range patterns {
		if count > maxRows {
			maxRows = count
		}
	}
	return maxRows
}

func main() {
	matrix := [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{1, 1, 0},
	}
	result := maxEqualRowsAfterFlips(matrix)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
