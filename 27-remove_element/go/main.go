package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for k := 0; k < len(nums); k++ {
		if nums[k] != val {
			nums[i] = nums[k]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Printf("Result: %#v\n", nums[:k])
	if fmt.Sprintf("%v", nums[:k]) == "[2 2]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
