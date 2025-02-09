package main

import (
	"fmt"
)

func countBadPairs(nums []int) int64 {
	differenceCounter := make(map[int64]int64)
	var badPairsCount int64
	for index, number := range nums {
		diff := int64(index - number)
		badPairsCount += int64(index) - differenceCounter[diff]
		differenceCounter[diff]++
	}
	return badPairsCount
}

func main() {
	nums := []int{4, 1, 3, 3}
	result := countBadPairs(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
