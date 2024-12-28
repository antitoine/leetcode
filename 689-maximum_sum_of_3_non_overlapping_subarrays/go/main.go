package main

import (
	"fmt"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	answer := make([]int, 3)

	currentSum := 0
	firstSubarraySum := 0
	secondSubarraySum := 0
	thirdSubarraySum := 0
	maxFirstSubarraySum := 0
	maxFirstSecondSum := 0

	maxFirstIndex := 0
	maxFirstSecondIndex1 := 0
	maxFirstSecondIndex2 := 0

	for i := k * 2; i < len(nums); i++ {
		firstSubarraySum += nums[i-k*2]
		secondSubarraySum += nums[i-k]
		thirdSubarraySum += nums[i]

		if i >= k*3-1 {
			if firstSubarraySum > maxFirstSubarraySum {
				maxFirstSubarraySum = firstSubarraySum
				maxFirstIndex = i - k*3 + 1
			}

			if maxFirstSubarraySum+secondSubarraySum > maxFirstSecondSum {
				maxFirstSecondSum = maxFirstSubarraySum + secondSubarraySum
				maxFirstSecondIndex1 = maxFirstIndex
				maxFirstSecondIndex2 = i - k*2 + 1
			}

			if maxFirstSecondSum+thirdSubarraySum > currentSum {
				currentSum = maxFirstSecondSum + thirdSubarraySum
				answer[0] = maxFirstSecondIndex1
				answer[1] = maxFirstSecondIndex2
				answer[2] = i - k + 1
			}

			firstSubarraySum -= nums[i-k*3+1]
			secondSubarraySum -= nums[i-k*2+1]
			thirdSubarraySum -= nums[i-k+1]
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
	k := 2
	result := maxSumOfThreeSubarrays(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[0 3 5]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
