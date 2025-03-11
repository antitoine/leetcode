package main

import (
	"fmt"
)

func numberOfSubstrings(s string) int {
	lastSeenPositions := [3]int{-1, -1, -1}
	substringCount := 0
	for index := 0; index < len(s); index++ {
		lastSeenPositions[s[index]-'a'] = index
		minLastSeenPosition := min(lastSeenPositions[0], lastSeenPositions[1], lastSeenPositions[2]) + 1
		substringCount += minLastSeenPosition
	}
	return substringCount
}

func main() {
	s := "abcabc"
	result := numberOfSubstrings(s)
	fmt.Printf("Result: %#v\n", result)
	if result == 10 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
