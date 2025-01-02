package main

import (
	"fmt"
	"strings"
)

const vowels = "aeiou"

func vowelString(word string) bool {
	return strings.IndexRune(vowels, rune(word[0])) != -1 && strings.IndexRune(vowels, rune(word[len(word)-1])) != -1
}

func vowelStrings(words []string, queries [][]int) []int {
	dp := make([]int, len(words)+1)

	for i := 1; i <= len(words); i++ {
		if vowelString(words[i-1]) {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = dp[query[1]+1] - dp[query[0]]
	}

	return result
}

func main() {
	words := []string{"aba", "bcb", "ece", "aa", "e"}
	queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
	result := vowelStrings(words, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[2 3 0]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
