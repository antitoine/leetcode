package main

import (
	"fmt"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(goal); i++ {
		if s[0] == goal[i] && s == (goal[i:]+goal[0:i]) {
			return true
		}
	}
	return false
}

func main() {
	s, goal := "abcde", "cdeab"
	result := rotateString(s, goal)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
