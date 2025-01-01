package main

import (
	"fmt"
	"strings"
)

func maxScore(s string) int {
	var result, zeros int
	ones := strings.Count(s, "1")

	for i := 0; i+1 < len(s); i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		if score := zeros + ones; score > result {
			result = score
		}
	}

	return result
}

func main() {
	s := "011101"
	result := maxScore(s)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
