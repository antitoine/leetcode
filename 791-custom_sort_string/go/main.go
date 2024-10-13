package main

import (
	"fmt"
)

func customSortString(order string, s string) string {
	characters := make(map[rune]int)
	for _, c := range s {
		characters[c]++
	}
	var result string
	for _, o := range order {
		if occ, found := characters[o]; found {
			for i := 0; i < occ; i++ {
				result += string(o)
			}
			delete(characters, o)
		}
	}
	for o, occ := range characters {
		for i := 0; i < occ; i++ {
			result += string(o)
		}
	}
	return result
}

func main() {
	input1, input2 := "cba", "abcd"
	result := customSortString(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == "cbad" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
