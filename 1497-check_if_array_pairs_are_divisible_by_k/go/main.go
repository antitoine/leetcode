package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {
	availableNums := make(map[int]int)
	for _, num := range arr {
		availableNums[num%k]++
	}
	for i := 1; i < k; i++ {
		smallOcc := availableNums[i]
		bigOcc := availableNums[k-i]
		negSmallOcc := availableNums[-i]
		negBigOcc := availableNums[i-k]
		if smallOcc-negSmallOcc != bigOcc-negBigOcc {
			return false
		}
	}
	return availableNums[0]%2 == 0
}

func main() {
	input1, input2 := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5
	result := canArrange(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == true {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
