package main

import (
	"fmt"
	"math"
)

func maxAbsoluteSum(nums []int) int {
	currentMaxSum, currentMinSum := 0, 0
	maxSum := 0
	for _, num := range nums {
		currentMaxSum = max(0, currentMaxSum) + num
		currentMinSum = min(0, currentMinSum) + num
		maxSum = max(maxSum, currentMaxSum, int(math.Abs(float64(currentMinSum))))
	}
	return maxSum
}

func main() {
	nums := []int{1, -3, 2, 3, -4}
	result := maxAbsoluteSum(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
