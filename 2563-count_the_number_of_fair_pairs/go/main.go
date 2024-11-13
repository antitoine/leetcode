package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var cnt int64
	for i := 0; i < len(nums)-1; i++ {
		left := sort.Search(len(nums)-i-1, func(j int) bool {
			return nums[i]+nums[i+j+1] >= lower
		}) + i + 1
		right := sort.Search(len(nums)-i-1, func(j int) bool {
			return nums[i]+nums[i+j+1] > upper
		}) + i
		cnt += int64(right - left + 1)
	}
	return cnt
}

func main() {
	nums, lower, upper := []int{0, 1, 7, 4, 4, 5}, 3, 6
	result := countFairPairs(nums, lower, upper)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
