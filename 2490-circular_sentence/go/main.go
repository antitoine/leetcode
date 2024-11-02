package main

import (
	"fmt"
	"strings"
)

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	lastWord := words[len(words)-1]
	if words[0][0] != lastWord[len(lastWord)-1] {
		return false
	}
	for i := 0; i+1 < len(words); i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}
	return true
}

func main() {
	sentence := "leetcode exercises sound delightful"
	result := isCircularSentence(sentence)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
