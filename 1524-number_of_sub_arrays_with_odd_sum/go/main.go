package main

import (
	"fmt"
)

const MOD = 1e9 + 7

func numOfSubarrays(arr []int) int {
	count := [2]int{1, 0}
	var answer, sum int
	for _, number := range arr {
		sum += number
		answer = (answer + count[(sum%2)^1]) % MOD
		count[sum%2]++
	}
	return answer
}

func main() {
	arr := []int{1, 3, 5}
	result := numOfSubarrays(arr)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
