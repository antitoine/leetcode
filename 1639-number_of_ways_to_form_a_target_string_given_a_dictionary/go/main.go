package main

import (
	"fmt"
)

func numWays(words []string, target string) int {
	const kMod int = 1_000_000_007
	wordLength := len(words[0])
	targetLength := len(target)

	dp := make([]int, targetLength+1)
	dp[0] = 1

	for j := 0; j < wordLength; j++ {
		count := make([]int, 26)
		for _, word := range words {
			count[word[j]-'a']++
		}
		for i := targetLength; i > 0; i-- {
			dp[i] = (dp[i] + dp[i-1]*count[target[i-1]-'a']) % kMod
		}
	}

	return dp[targetLength]
}

func main() {
	words := []string{"acca", "bbbb", "caca"}
	target := "aba"
	result := numWays(words, target)
	fmt.Printf("Result: %#v\n", result)
	if result == 6 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
