package main

import (
	"fmt"
)

func takeCharacters(s string, k int) int {
	charCounts := [3]int{}
	for _, c := range s {
		charCounts[c-'a']++
	}
	if charCounts[0] < k || charCounts[1] < k || charCounts[2] < k {
		return -1
	}
	maxSubstringLength := 0
	windowStart := 0
	stringSize := len(s)
	for windowEnd := 0; windowEnd < stringSize; windowEnd++ {
		currentCharIndex := s[windowEnd] - 'a'
		charCounts[currentCharIndex]--

		for charCounts[currentCharIndex] < k {
			charCounts[s[windowStart]-'a']++
			windowStart++
		}

		if maxSubstringLength < windowEnd-windowStart+1 {
			maxSubstringLength = windowEnd - windowStart + 1
		}
	}
	return stringSize - maxSubstringLength
}

func main() {
	s, k := "aabaaaacaabc", 2
	result := takeCharacters(s, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 8 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
