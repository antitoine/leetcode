package main

import (
	"fmt"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	n, m := len(robot), len(factory)
	sort.Ints(robot)                                                                  // Sort robots by position
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] }) // Sort factories by position

	// Initialize a DP table with large values (infinity)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
		for j := range dp[i] {
			dp[i][j] = 9007199254740991
		}
	}

	// Base case: zero robots means zero distance
	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}

	// Fill DP table
	for j := 1; j <= m; j++ {
		posJ, limitJ := int64(factory[j-1][0]), factory[j-1][1]
		for i := 1; i <= n; i++ {
			dp[i][j] = dp[i][j-1] // Option to skip this factory
			sumDistance := int64(0)
			for k := 1; k <= limitJ && k <= i; k++ {
				sumDistance += abs(int64(robot[i-k]) - posJ) // Cumulative distance for k robots
				dp[i][j] = min(dp[i][j], dp[i-k][j-1]+sumDistance)
			}
		}
	}

	return dp[n][m]
}

// abs calculates the absolute value of an int64
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	robot := []int{0, 4, 6}
	factory := [][]int{{2, 2}, {6, 2}}
	result := minimumTotalDistance(robot, factory)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
