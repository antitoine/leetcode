package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	i := 0
	for k := 0; k < len(nums); k++ {
		if k == 0 || nums[k-1] < nums[k] {
			nums[i] = nums[k]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Printf("Result: %#v\n", nums[:k])
	if fmt.Sprintf("%v", nums[:k]) == "[1 2]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
