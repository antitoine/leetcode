package main

import (
	"fmt"
	"strings"
)

func countPrefixSuffixPairs(words []string) int {
	var result int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				result++
			}
		}
	}
	return result
}

func main() {
	words := []string{"a", "aba", "ababa", "aa"}
	result := countPrefixSuffixPairs(words)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
