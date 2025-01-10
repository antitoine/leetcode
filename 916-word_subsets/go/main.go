package main

import (
	"fmt"
)

func wordSubsets(words1 []string, words2 []string) []string {
	words2MaxOcc := make(map[rune]int)
	for _, word := range words2 {
		wordOcc := make(map[rune]int)
		for _, c := range word {
			wordOcc[c]++
		}
		for c, occ := range wordOcc {
			if maxOcc, found := words2MaxOcc[c]; !found || maxOcc < occ {
				words2MaxOcc[c] = occ
			}
		}
	}

	var result []string
	for _, word := range words1 {
		wordOcc := make(map[rune]int)
		for _, c := range word {
			wordOcc[c]++
		}
		isUniversal := true
		for requiredChar, requiredOcc := range words2MaxOcc {
			if occ, found := wordOcc[requiredChar]; !found || occ < requiredOcc {
				isUniversal = false
				break
			}
		}
		if isUniversal {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}
	result := wordSubsets(words1, words2)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[facebook google leetcode]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
