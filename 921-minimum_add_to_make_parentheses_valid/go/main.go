package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	balance := 0
	op := 0
	for _, char := range s {
		if char == '(' {
			balance++
		} else if balance > 0 {
			balance--
		} else {
			op++
		}
	}
	if balance >= 0 {
		return balance + op
	} else {
		return -balance + op
	}
}

func main() {
	input := "())"
	result := minAddToMakeValid(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
