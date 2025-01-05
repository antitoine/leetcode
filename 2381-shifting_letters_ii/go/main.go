package main

import (
	"fmt"
)

func shiftingLetters(s string, shifts [][]int) string {
	delta := make([]int, len(s)+1)
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if shift[2] == 0 {
			direction = -1
		}
		delta[start] += direction
		delta[end+1] -= direction
	}
	for i := 1; i <= len(s); i++ {
		delta[i] += delta[i-1]
	}

	result := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		newChar := (int(s[i]-'a') + (delta[i]%26 + 26)) % 26
		result[i] = rune('a' + newChar)
	}
	return string(result)
}

func main() {
	s := "abc"
	shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}
	result := shiftingLetters(s, shifts)
	fmt.Printf("Result: %#v\n", result)
	if result == "ace" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
