package main

import (
	"fmt"
)

func isArraySpecial(nums []int) bool {
	previousIsPair := nums[0]%2 == 0
	for i := 1; i < len(nums); i++ {
		currentIsPair := nums[i]%2 == 0
		if previousIsPair == currentIsPair {
			return false
		}
		previousIsPair = currentIsPair
	}
	return true
}

func main() {
	nums := []int{2, 1, 4}
	result := isArraySpecial(nums)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
