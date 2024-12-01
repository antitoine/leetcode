package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	values := make(map[int]struct{})
	for _, v := range arr {
		if _, found := values[v]; found {
			if v == 0 {
				return true
			}
			continue
		}
		values[v] = struct{}{}
		if v == 0 {
			continue
		}
		if _, found := values[v*2]; found {
			return true
		}
		if v%2 == 0 {
			if _, found := values[v/2]; found {
				return true
			}
		}
	}
	return false
}

func main() {
	input := []int{10, 2, 5, 3}
	result := checkIfExist(input)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
