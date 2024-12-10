package main

import (
	"fmt"
)

func maximumLength(s string) int {
	n := len(s)
	ans := -1
	runningLen := 0
	prevLetter := '@'

	counts := make([][]int, 26)
	for i := 0; i < 26; i++ {
		counts[i] = make([]int, n+1)
	}

	for _, c := range s {
		if c == prevLetter {
			runningLen++
			counts[c-'a'][runningLen]++
		} else {
			runningLen = 1
			counts[c-'a'][runningLen]++
			prevLetter = c
		}
	}

	for _, count := range counts {
		ans = max(ans, getMaxFreq(count, n))
	}

	return ans
}

func getMaxFreq(count []int, maxFreq int) int {
	times := 0
	for freq := maxFreq; freq >= 1; freq-- {
		times += count[freq]
		if times >= 3 {
			return freq
		}
	}
	return -1
}

func main() {
	input := "aaaa"
	result := maximumLength(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
