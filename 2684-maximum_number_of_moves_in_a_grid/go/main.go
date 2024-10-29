package main

import (
	"fmt"
)

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 // Initialize with -1 to mark unvisited cells
		}
	}

	// Define the possible directions (down-right, right, up-right)
	directions := [][2]int{{-1, 1}, {0, 1}, {1, 1}}

	// Helper function to perform DFS with memoization
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		// If already computed, return the stored result
		if dp[row][col] != -1 {
			return dp[row][col]
		}

		maxMove := 0
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			// Check if within bounds and if the next cell has a strictly larger value
			if newRow >= 0 && newRow < m && newCol < n && grid[newRow][newCol] > grid[row][col] {
				maxMove = max(maxMove, 1+dfs(newRow, newCol))
			}
		}

		dp[row][col] = maxMove // Memoize the result
		return dp[row][col]
	}

	// Start DFS from each cell in the first column and find the maximum moves
	maxM := 0
	for r := 0; r < m; r++ {
		maxM = max(maxM, dfs(r, 0))
	}

	return maxM
}

func main() {
	input := [][]int{
		{2, 4, 3, 5},
		{5, 4, 9, 3},
		{3, 4, 2, 11},
		{10, 9, 13, 15},
	}
	result := maxMoves(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
