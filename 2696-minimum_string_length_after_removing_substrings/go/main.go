package main

import (
	"fmt"
	"strings"
)

func minLength(s string) int {
	if strings.Contains(s, "AB") {
		return minLength(strings.ReplaceAll(s, "AB", ""))
	} else if strings.Contains(s, "CD") {
		return minLength(strings.ReplaceAll(s, "CD", ""))
	} else {
		return len(s)
	}
}

func main() {
	input := "ABFCACDB"
	result := minLength(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
