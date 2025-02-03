package main

import (
	"fmt"
)

func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := 1
	increasing := 1
	decreasing := 1

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			increasing++
			decreasing = 1
		} else if nums[i] < nums[i-1] {
			decreasing++
			increasing = 1
		} else {
			increasing = 1
			decreasing = 1
		}
		ans = max(ans, max(increasing, decreasing))
	}

	return ans
}

func main() {
	nums := []int{1, 4, 3, 3, 2}
	result := longestMonotonicSubarray(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
