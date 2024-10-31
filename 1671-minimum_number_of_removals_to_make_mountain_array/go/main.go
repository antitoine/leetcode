package main

import (
	"fmt"
)

func minimumMountainRemovals(nums []int) int {
	n := len(nums)

	// Step 1: Calculate LIS ending at each index
	LIS := make([]int, n)
	for i := range LIS {
		LIS[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if LIS[i] < LIS[j]+1 {
					LIS[i] = LIS[j] + 1
				}
			}
		}
	}

	// Step 2: Calculate LDS starting from each index
	LDS := make([]int, n)
	for i := range LDS {
		LDS[i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				if LDS[i] < LDS[j]+1 {
					LDS[i] = LDS[j] + 1
				}
			}
		}
	}

	// Step 3: Find the longest mountain length
	maxMountainLength := 0
	for i := 1; i < n-1; i++ {
		if LIS[i] > 1 && LDS[i] > 1 {
			mountainLength := LIS[i] + LDS[i] - 1
			if mountainLength > maxMountainLength {
				maxMountainLength = mountainLength
			}
		}
	}

	// Step 4: Minimum elements to remove
	return n - maxMountainLength
}

func main() {
	input := []int{2, 1, 1, 5, 6, 2, 3, 1}
	result := minimumMountainRemovals(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
