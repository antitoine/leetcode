package main

import (
	"fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	swappedIdx := -1
	swapped := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if swappedIdx == -1 {
				swappedIdx = i
			} else if !swapped {
				if s1[swappedIdx] != s2[i] || s1[i] != s2[swappedIdx] {
					return false
				}
				swapped = true
			} else {
				return false
			}
		}
	}
	return swappedIdx == -1 || swapped
}

func main() {
	s1, s2 := "bank", "kanb"
	result := areAlmostEqual(s1, s2)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
