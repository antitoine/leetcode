package main

import (
	"fmt"
)

func maxChunksToSorted(arr []int) int {
	chunkCount := 0
	maxElement := 0

	for index, value := range arr {
		if value > maxElement {
			maxElement = value
		}
		if index == maxElement {
			chunkCount++
		}
	}

	return chunkCount
}

func main() {
	arr := []int{4, 3, 2, 1, 0}
	result := maxChunksToSorted(arr)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
