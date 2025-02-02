package main

import (
	"fmt"
)

func check(nums []int) bool {
	rotated := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if rotated {
				return false
			}
			rotated = true
		}
	}
	if rotated && nums[0] < nums[len(nums)-1] {
		return false
	}
	return true
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	result := check(nums)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
