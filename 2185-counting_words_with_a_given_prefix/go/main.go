package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	var result int
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			result++
		}
	}
	return result
}

func main() {
	words := []string{"pay", "**at**tention", "practice", "**at**tend"}
	pref := "at"
	result := prefixCount(words, pref)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
