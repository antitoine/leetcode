package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	remainder := sum % p
	if remainder == 0 {
		return 0
	}

	prefixToIndex := make(map[int]int)
	prefixToIndex[0] = -1
	ans := len(nums)
	prefix := 0

	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		prefix %= p
		target := (prefix - remainder + p) % p
		if idx, found := prefixToIndex[target]; found {
			if i-idx < ans {
				ans = i - idx
			}
		}
		prefixToIndex[prefix] = i
	}

	if ans == len(nums) {
		return -1
	}

	return ans
}

func main() {
	input1, input2 := []int{3, 1, 4, 2}, 6
	result := minSubarray(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
