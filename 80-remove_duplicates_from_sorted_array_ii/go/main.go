package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	j := min(2, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[j-2] == nums[i] && nums[j-1] == nums[i] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Printf("Result: %#v\n", nums[:k])
	if fmt.Sprintf("%v", nums[:k]) == "[1 1 2 2 3]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
