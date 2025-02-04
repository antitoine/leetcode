package main

import (
	"fmt"
)

func maxAscendingSum(nums []int) int {
	var maxSum int
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			currentSum += nums[i]
		} else {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			currentSum = nums[i]
		}
	}
	if currentSum > maxSum {
		maxSum = currentSum
	}
	return maxSum
}

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	result := maxAscendingSum(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 65 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
