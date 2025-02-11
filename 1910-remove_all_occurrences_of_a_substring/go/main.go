package main

import (
	"fmt"
	"strings"
)

func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		index := strings.Index(s, part)
		s = s[:index] + s[index+len(part):]
	}
	return s
}

func main() {
	s := "daabcbaabcbc"
	part := "abc"
	result := removeOccurrences(s, part)
	fmt.Printf("Result: %#v\n", result)
	if result == "dab" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
