package main

import (
	"fmt"
)

func clearDigits(s string) string {
	var result string
	var toRemove int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			toRemove++
		} else if toRemove == 0 {
			result = string(s[i]) + result
		} else {
			toRemove--
		}
	}
	return result
}

func main() {
	s := "abc"
	result := clearDigits(s)
	fmt.Printf("Result: %#v\n", result)
	if result == "abc" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
