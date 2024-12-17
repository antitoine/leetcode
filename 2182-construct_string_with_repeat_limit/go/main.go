package main

import (
	"fmt"
	"strings"
)

func repeatLimitedString(s string, repeatLimit int) string {
	var ans strings.Builder
	count := make(map[rune]int)

	for _, c := range s {
		count[c]++
	}

	for {
		addOne := ans.Len() > 0 && shouldAddOne(ans.String(), count)
		c := getLargestChar(ans.String(), count)
		if c == ' ' {
			break
		}
		repeats := 1
		if !addOne {
			repeats = min(count[c], repeatLimit)
		}
		ans.WriteString(strings.Repeat(string(c), repeats))
		count[c] -= repeats
	}

	return ans.String()
}

func shouldAddOne(ans string, count map[rune]int) bool {
	for c := 'z'; c >= 'a'; c-- {
		if count[c] > 0 {
			return rune(ans[len(ans)-1]) == c
		}
	}
	return false
}

func getLargestChar(ans string, count map[rune]int) rune {
	for c := 'z'; c >= 'a'; c-- {
		if count[c] > 0 && (len(ans) == 0 || rune(ans[len(ans)-1]) != c) {
			return c
		}
	}
	return ' '
}

func main() {
	s := "cczazcc"
	repeatLimit := 3
	result := repeatLimitedString(s, repeatLimit)
	fmt.Printf("Result: %#v\n", result)
	if result == "zzcccac" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
