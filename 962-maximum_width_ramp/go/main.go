package main

import (
	"fmt"
)

func maxWidthRamp(nums []int) int {
	startingIdx := 1
	for ; startingIdx < len(nums) && nums[startingIdx-1] > nums[startingIdx]; startingIdx++ {
	}
	if startingIdx == len(nums) {
		return 0
	}
	for startIdx := 0; startIdx < len(nums); startIdx++ {
		for idx := 0; idx <= startIdx; idx++ {
			endIdx := idx + len(nums) - startIdx - 1
			if nums[idx] <= nums[endIdx] {
				return endIdx - idx
			}
		}
	}
	return 0
}

func main() {
	input := []int{6, 0, 8, 2, 1, 5}
	result := maxWidthRamp(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
