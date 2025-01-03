package main

import (
	"fmt"
)

func waysToSplitArray(nums []int) int {
	leftSum := make([]int, len(nums))
	leftSum[0] = nums[0]
	rightSum := make([]int, len(nums))
	rightSum[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
		rightSum[len(nums)-1-i] = rightSum[len(nums)-i] + nums[len(nums)-1-i]
	}

	var result int
	for i := 0; i < len(nums)-1; i++ {
		if leftSum[i] >= rightSum[i+1] {
			result++
		}
	}

	return result
}

func main() {
	nums := []int{10, 4, -8, 7}
	result := waysToSplitArray(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
