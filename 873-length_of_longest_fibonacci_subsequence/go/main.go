package main

import (
	"fmt"
)

func lenLongestFibSubseq(arr []int) int {
	indexMap := make(map[int]int)
	n := len(arr)
	for i, num := range arr {
		indexMap[num] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 2
		}
	}
	maxLength := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			prevNum := arr[i] - arr[j]
			if k, exists := indexMap[prevNum]; exists && k < j {
				dp[j][i] = max(dp[j][i], dp[k][j]+1)
				maxLength = max(maxLength, dp[j][i])
			}
		}
	}
	if maxLength > 2 {
		return maxLength
	}
	return 0
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := lenLongestFibSubseq(arr)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
