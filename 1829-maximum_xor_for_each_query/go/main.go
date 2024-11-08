package main

import (
	"fmt"
	"math"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	result := make([]int, len(nums))
	result[len(nums)-1] = nums[0]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] ^ nums[len(nums)-i-1]
	}
	maxInt := int(math.Pow(float64(2), float64(maximumBit))) - 1
	for i := 0; i < len(nums); i++ {
		result[i] = ^result[i] & maxInt
	}
	return result
}

func main() {
	nums := []int{0, 1, 1, 3}
	maximumBit := 2
	result := getMaximumXor(nums, maximumBit)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[0 3 2 3]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
