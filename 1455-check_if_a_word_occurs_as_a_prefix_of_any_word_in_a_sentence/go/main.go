package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, word := range strings.Split(sentence, " ") {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}

func main() {
	sentence, searchWord := "i love eating burger", "burg"
	result := isPrefixOfWord(sentence, searchWord)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
