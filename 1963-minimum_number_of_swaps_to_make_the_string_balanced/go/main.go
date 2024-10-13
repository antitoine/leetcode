package main

import (
	"fmt"
)

func minSwaps(s string) int {
	var stack []rune
	remainingBrackets := 0

	for _, current := range s {
		if current == '[' {
			stack = append(stack, current)
		} else if current == ']' && len(stack) == 0 {
			remainingBrackets++
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return (remainingBrackets + 1) / 2
}

func main() {
	input := "][]["
	result := minSwaps(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
