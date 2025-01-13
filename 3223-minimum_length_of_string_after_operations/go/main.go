package main

import (
	"fmt"
)

func minimumLength(s string) int {
	charOcc := make(map[rune]int)
	var finalLength int
	for _, c := range s {
		charOcc[c]++
		finalLength++
		if charOcc[c] == 3 {
			charOcc[c] -= 2
			finalLength -= 2
		}
	}
	return finalLength
}

func main() {
	s := "abaacbcbb"
	result := minimumLength(s)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
