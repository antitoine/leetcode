package main

import (
	"fmt"
)

func tupleSameProduct(nums []int) int {
	multiples := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			multiples[nums[i]*nums[j]]++
		}
	}
	var result int
	for _, v := range multiples {
		if v > 1 {
			result += v * (v - 1) * 4
		}
	}
	return result
}

func main() {
	input := []int{2, 3, 4, 6}
	result := tupleSameProduct(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 8 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
