package main

import (
	"fmt"
)

func nextCycle(r byte) byte {
	if r == 'z' {
		return 'a'
	}
	return r + 1
}

func canMakeSubsequence(str1 string, str2 string) bool {
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] || nextCycle(str1[i]) == str2[j] {
			j++
		}
		i++
	}
	return j == len(str2)
}

func main() {
	str1, str2 := "abc", "ad"
	result := canMakeSubsequence(str1, str2)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
