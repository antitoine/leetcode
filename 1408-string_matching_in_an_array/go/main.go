package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	result := stringMatching(words)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == `[as hero]` {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
