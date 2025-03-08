package main

import (
	"fmt"
	"strings"
)

func minimumRecolors(blocks string, k int) int {
	whiteBlockCount := strings.Count(blocks[:k], "W")
	minRecolors := whiteBlockCount
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whiteBlockCount++
		}
		if blocks[i-k] == 'W' {
			whiteBlockCount--
		}
		if whiteBlockCount < minRecolors {
			minRecolors = whiteBlockCount
		}
	}
	return minRecolors
}

func main() {
	blocks := "WBBWWBBWBW"
	k := 7
	result := minimumRecolors(blocks, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
