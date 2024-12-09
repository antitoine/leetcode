package main

import (
	"fmt"
	"slices"
)

func isArraySpecial(nums []int, queries [][]int) []bool {
	var invalidPairs []int
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			invalidPairs = append(invalidPairs, i)
		}
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		_, found := slices.BinarySearchFunc(invalidPairs, query, func(pair int, query []int) int {
			if pair > query[1] {
				return 1
			} else if pair-1 < query[0] {
				return -1
			}
			return 0
		})
		result[i] = !found
	}

	return result
}

func main() {
	nums, queries := []int{3, 4, 1, 2, 6}, [][]int{{0, 4}}
	result := isArraySpecial(nums, queries)
	fmt.Printf("Result: %#v\n", result)
	if !result[0] {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
