package main

import (
	"fmt"
	"math"
)

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	ans := int64(math.MaxInt64)
	var sumRow0, sumRow1 int64

	for _, val := range grid[0] {
		sumRow0 += int64(val)
	}

	for i := 0; i < n; i++ {
		sumRow0 -= int64(grid[0][i])
		ans = min(ans, max(sumRow0, sumRow1))
		sumRow1 += int64(grid[1][i])
	}

	return ans
}

func main() {
	grid := [][]int{
		{2, 5, 4},
		{1, 5, 1},
	}
	result := gridGame(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
