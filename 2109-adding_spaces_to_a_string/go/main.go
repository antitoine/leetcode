package main

import (
	"fmt"
)

func addSpaces(s string, spaces []int) string {
	result := make([]byte, len(s)+len(spaces))
	k := 0
	for i := 0; i < len(s); i++ {
		if k < len(spaces) && spaces[k] == i {
			result[i+k] = ' '
			k++
		}
		result[i+k] = s[i]
	}
	return string(result)
}

func main() {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}
	result := addSpaces(s, spaces)
	fmt.Printf("Result: %#v\n", result)
	if result == "Leetcode Helps Me Learn" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
