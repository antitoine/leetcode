package main

import (
	"fmt"
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	maxLength := 0
	i, j := 0, 1
	for j < len(nums) {
		if nums[i]+k < nums[j]-k {
			if maxLength < j-i {
				maxLength = j - i
			}
			i++
		}
		j++
	}
	if maxLength < j-i {
		maxLength = j - i
	}
	return maxLength
}

func main() {
	nums := []int{4, 6, 1, 2}
	k := 2
	result := maximumBeauty(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
