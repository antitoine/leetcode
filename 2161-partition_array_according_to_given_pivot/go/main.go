package main

import (
	"fmt"
)

func pivotArray(nums []int, pivot int) []int {
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)
	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}
	return append(less, append(equal, greater...)...)
}

func main() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	result := pivotArray(nums, pivot)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[9 5 3 10 10 12 14]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
