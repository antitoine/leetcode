package main

import (
	"fmt"
)

func makeFancyString(s string) string {
	if len(s) < 2 {
		return s
	}
	result := make([]rune, len(s))
	result[0] = rune(s[0])
	result[1] = rune(s[1])
	resultIdx := 2
	for sIdx := 2; sIdx < len(s); sIdx++ {
		if !(result[resultIdx-1] == rune(s[sIdx]) && result[resultIdx-2] == rune(s[sIdx])) {
			result[resultIdx] = rune(s[sIdx])
			resultIdx++
		}
	}
	return string(result[0:resultIdx])
}

func main() {
	input := "leeetcode"
	result := makeFancyString(input)
	fmt.Printf("Result: %#v\n", result)
	if result == "leetcode" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
