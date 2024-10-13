package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1Freq := getCharFreq(s1)
	for s2StartIdx := 0; s2StartIdx+len(s1) <= len(s2); s2StartIdx++ {
		s2SubStr := s2[s2StartIdx : s2StartIdx+len(s1)]
		s2SubStrFreq := getCharFreq(s2SubStr)
		if charFreqEquals(s1Freq, s2SubStrFreq) {
			return true
		}
	}
	return false
}

func getCharFreq(str string) []int {
	freq := make([]int, 26)
	for _, char := range str {
		freq[char-97] += 1
	}
	return freq
}

func charFreqEquals(s1Freq []int, s2Freq []int) bool {
	for idx, freq := range s1Freq {
		if s2Freq[idx] != freq {
			return false
		}
	}
	return true
}

func main() {
	input1, input2 := "ab", "eidbaooo"
	result := checkInclusion(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == true {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
