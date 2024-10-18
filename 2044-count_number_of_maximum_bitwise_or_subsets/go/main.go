package main

import (
	"fmt"
)

func findSubsets(nums []int, size, bitwiseOr, index, currSize, currOr int, currentSubset []int) int {
	if currSize > size {
		return 0
	}

	if currSize == size {
		if currOr == bitwiseOr {
			return 1
		}
		return 0
	}

	if index == len(nums) {
		return 0
	}

	var count int
	count += findSubsets(nums, size, bitwiseOr, index+1, currSize, currOr, currentSubset)

	newOr := currOr | nums[index]
	currentSubset = append(currentSubset, nums[index])
	count += findSubsets(nums, size, bitwiseOr, index+1, currSize+1, newOr, currentSubset)

	return count
}

func getAllSubOfFixedSizeAndBitwiseOr(nums []int, size int, bitwiseOr int) int {
	return findSubsets(nums, size, bitwiseOr, 0, 0, 0, []int{})
}

func countMaxOrSubsets(nums []int) int {
	maxBitwiseOr := 0
	for _, num := range nums {
		maxBitwiseOr |= num
	}
	nbGrp := 0
	for grpSize := 1; grpSize <= len(nums); grpSize++ {
		nbGrp += getAllSubOfFixedSizeAndBitwiseOr(nums, grpSize, maxBitwiseOr)
	}
	return nbGrp
}

func main() {
	input := []int{3, 2, 1, 5}
	result := countMaxOrSubsets(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 6 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
