package main

import (
	"fmt"
)

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	charOcc := make(map[rune]int)
	for _, c := range s {
		charOcc[c]++
	}
	var nbOddOcc int
	for _, occ := range charOcc {
		if occ%2 == 1 {
			nbOddOcc++
		}
	}
	if nbOddOcc > k {
		return false
	}
	return true
}

func main() {
	s := "annabelle"
	k := 2
	result := canConstruct(s, k)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
