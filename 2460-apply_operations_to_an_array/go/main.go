package main

import (
	"fmt"
)

func applyOperations(nums []int) []int {
	result := make([]int, len(nums))
	in := 0
	out := 0
	for in < len(nums) {
		if in < len(nums)-1 && nums[in] == nums[in+1] {
			nums[in+1] = 0
			if nums[in] > 0 {
				result[out] = nums[in] * 2
				out++
			}
			in++
		} else if nums[in] == 0 {
			in++
		} else {
			result[out] = nums[in]
			in++
			out++
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 2, 1, 1, 0}
	result := applyOperations(nums)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[1 4 2 0 0 0]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
