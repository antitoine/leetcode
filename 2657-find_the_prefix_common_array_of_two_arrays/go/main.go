package main

import (
	"fmt"
)

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	result := make([]int, n)
	occ := make([]int, n)
	k := 0
	for i := 0; i < n; i++ {
		occ[A[i]-1]++
		if occ[A[i]-1] == 2 {
			k++
		}
		occ[B[i]-1]++
		if occ[B[i]-1] == 2 {
			k++
		}
		result[i] = k
	}
	return result
}

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	result := findThePrefixCommonArray(A, B)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[0 2 3 4]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
