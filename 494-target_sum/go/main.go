package main

import (
	"fmt"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum < abs(target) || (sum+target)%2 == 1 {
		return 0
	}

	t := (sum + target) / 2
	dp := make([]int, t+1)
	dp[0] = 1
	for _, num := range nums {
		for i := t; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[t]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	result := findTargetSumWays(nums, target)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
