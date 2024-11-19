package main

import (
	"fmt"
)

func maximumSubarraySum(nums []int, k int) int64 {
	elements := make(map[int]struct{})
	var currentSum int64
	var maxSum int64
	for i, j := 0, 0; j < len(nums); j++ {
		_, found := elements[nums[j]]
		for found {
			currentSum -= int64(nums[i])
			delete(elements, nums[i])
			i++
			_, found = elements[nums[j]]
		}
		currentSum += int64(nums[j])
		elements[nums[j]] = struct{}{}
		if j-i == k-1 {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			currentSum -= int64(nums[i])
			delete(elements, nums[i])
			i++
		}
	}
	return maxSum
}

func main() {
	nums, k := []int{1, 5, 4, 2, 9, 9, 9}, 3
	result := maximumSubarraySum(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 15 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
