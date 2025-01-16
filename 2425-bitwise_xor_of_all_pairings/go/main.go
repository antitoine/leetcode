package main

import (
	"fmt"
)

func xorAllNums(nums1 []int, nums2 []int) int {
	resultXor := 0
	if len(nums2)%2 != 0 {
		for _, num := range nums1 {
			resultXor ^= num
		}
	}
	if len(nums1)%2 != 0 {
		for _, num := range nums2 {
			resultXor ^= num
		}
	}
	return resultXor
}

func main() {
	nums1 := []int{2, 1, 3}
	nums2 := []int{10, 2, 5, 0}
	result := xorAllNums(nums1, nums2)
	fmt.Printf("Result: %#v\n", result)
	if result == 13 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
