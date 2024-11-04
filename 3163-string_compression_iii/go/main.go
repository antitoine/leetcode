package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressedString(word string) string {
	if len(word) == 0 {
		return word
	}
	var comp strings.Builder
	lastCharFreq := 1
	for i := 1; i < len(word); i++ {
		if lastCharFreq < 9 && word[i-1] == word[i] {
			lastCharFreq++
		} else {
			comp.WriteString(strconv.Itoa(lastCharFreq))
			comp.WriteByte(word[i-1])
			lastCharFreq = 1
		}
	}
	comp.WriteString(strconv.Itoa(lastCharFreq))
	comp.WriteByte(word[len(word)-1])
	return comp.String()
}

func main() {
	input := "aaaaaaaaaaaaaabb"
	result := compressedString(input)
	fmt.Printf("Result: %#v\n", result)
	if result == "9a5a2b" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
