package main

import (
	"fmt"
)

func countSquares(matrix [][]int) int {
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				matrix[i][j] = 1 + min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1]))
			}
			result += matrix[i][j]
		}
	}
	return result
}

func main() {
	input := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	result := countSquares(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 15 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
