package main

import (
	"fmt"
	"math"
	"slices"
)

func isMaxBallsPossible(nums []int, maxOperations, maxBalls int) bool {
	var ops int
	for _, num := range nums {
		ops += int(math.Ceil(float64(num)/float64(maxBalls)) - 1)
		if ops > maxOperations {
			return false
		}
	}
	return true
}

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, slices.Max(nums)
	for left < right {
		middle := left + ((right - left) / 2)
		if isMaxBallsPossible(nums, maxOperations, middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func main() {
	nums := []int{9}
	maxOperations := 2
	result := minimumSize(nums, maxOperations)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
