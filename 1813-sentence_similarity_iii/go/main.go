package main

import (
	"fmt"
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}

	sentence1Words := strings.Split(sentence1, " ")
	sentence2Words := strings.Split(sentence2, " ")

	startIdx1 := 0
	startIdx2 := 0
	for startIdx1 < len(sentence1Words) && startIdx2 < len(sentence2Words) && sentence1Words[startIdx1] == sentence2Words[startIdx2] {
		startIdx1++
		startIdx2++
	}

	if startIdx1 == len(sentence1Words) {
		return startIdx1 < len(sentence2Words)
	} else if startIdx2 == len(sentence2Words) {
		return startIdx2 < len(sentence1Words)
	}

	endIdx1 := len(sentence1Words) - 1
	endIdx2 := len(sentence2Words) - 1
	for endIdx1 > startIdx1 && endIdx2 > startIdx2 && sentence1Words[endIdx1] == sentence2Words[endIdx2] {
		endIdx1--
		endIdx2--
	}

	var sentenceUpdated string
	var sentenceRef string
	if startIdx1 == endIdx1 {
		sentenceUpdated = strings.Join(append(sentence1Words[:startIdx1], append(sentence2Words[startIdx2:endIdx2], sentence1Words[endIdx1:]...)...), " ")
		sentenceRef = sentence2
	} else if startIdx2 == endIdx2 {
		sentenceUpdated = strings.Join(append(sentence2Words[:startIdx2], append(sentence1Words[startIdx1:endIdx1], sentence2Words[endIdx2:]...)...), " ")
		sentenceRef = sentence1
	} else {
		return false
	}

	return sentenceUpdated == sentenceRef
}

func main() {
	input1, input2 := "My name is Haley", "My Haley"
	result := areSentencesSimilar(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == true {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
