package main

import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {
	var result int
	first := make([]int, 26)
	last := make([]int, 26)

	for i := 0; i < 26; i++ {
		first[i] = len(s)
	}

	for i, c := range s {
		index := int(c - 'a')
		if i < first[index] {
			first[index] = i
		}
		last[index] = i
	}

	for i := 0; i < 26; i++ {
		f, l := first[i], last[i]
		if f < l {
			charSet := make(map[byte]struct{})
			for j := f + 1; j < l; j++ {
				charSet[s[j]] = struct{}{}
			}
			result += len(charSet)
		}
	}

	return result
}

func main() {
	s := "aabca"
	result := countPalindromicSubsequence(s)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
