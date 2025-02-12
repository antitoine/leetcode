package main

import (
	"fmt"
)

func maximumSum(nums []int) int {
	maxSum := -1
	digitSumMaxNum := make(map[int]int)
	for _, num := range nums {
		digitSum := 0
		tempNum := num
		for tempNum > 0 {
			digitSum += tempNum % 10
			tempNum /= 10
		}
		if prevNum, exists := digitSumMaxNum[digitSum]; exists {
			maxSum = max(maxSum, prevNum+num)
		}
		digitSumMaxNum[digitSum] = max(digitSumMaxNum[digitSum], num)
	}
	return maxSum
}

func main() {
	nums := []int{18, 43, 36, 13, 7}
	result := maximumSum(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 54 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
