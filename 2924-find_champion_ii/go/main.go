package main

import (
	"fmt"
)

func minChanges(s string) int {
	nbChanges := 0
	for i := 0; i+1 < len(s); i += 2 {
		if s[i] != s[i+1] {
			nbChanges++
		}
	}
	return nbChanges
}

func main() {
	input := "1001"
	result := minChanges(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
