package main

import (
	"fmt"
)

func computeS(previousS []bool) []bool {
	newSize := 2*len(previousS) + 1
	result := make([]bool, newSize)
	for i := 0; i < len(previousS); i++ {
		result[i] = previousS[i]
		result[newSize-i-1] = !previousS[i]
	}
	result[len(previousS)] = true
	return result
}

func findKthBit(n int, k int) byte {
	s := []bool{false}
	for i := 2; i <= n; i++ {
		s = computeS(s)
	}
	if s[k-1] {
		return '1'
	}
	return '0'
}

func main() {
	input1, input2 := 3, 1
	result := findKthBit(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == '0' {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
