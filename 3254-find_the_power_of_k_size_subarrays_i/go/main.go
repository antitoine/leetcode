package main

import (
	"fmt"
)

func resultsArray(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	i := 0
	for i < len(result) {
		j := i + 1
		for j < i+k {
			if nums[j-1]+1 != nums[j] {
				for i < j && i < len(result) {
					result[i] = -1
					i++
				}
				break
			}
			j++
		}
		if j == i+k {
			result[i] = nums[j-1]
			i++
			for i < len(result) && nums[j-1]+1 == nums[j] {
				result[i] = nums[j]
				i++
				j++
			}
		}
	}
	return result
}
func main() {
	nums, k := []int{1, 2, 3, 4, 3, 2, 5}, 3
	result := resultsArray(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[3 4 -1 -1 -1]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
