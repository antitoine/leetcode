package main

import (
	"fmt"
)

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}
	if left == n-1 {
		return 0
	}
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	minSubarrayLength := min(n-left-1, right)
	for l, r := 0, right; l <= left; l++ {
		for r < n && arr[r] < arr[l] {
			r++
		}
		minSubarrayLength = min(minSubarrayLength, r-l-1)
	}
	return minSubarrayLength
}

func main() {
	arr := []int{1, 2, 3, 10, 4, 2, 3, 5}
	result := findLengthOfShortestSubarray(arr)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
