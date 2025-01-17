package main

import (
	"fmt"
)

func doesValidArrayExist(derived []int) bool {
	var result int
	for _, val := range derived {
		result ^= val
	}
	return result == 0
}

func main() {
	derived := []int{1, 1, 0}
	result := doesValidArrayExist(derived)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
