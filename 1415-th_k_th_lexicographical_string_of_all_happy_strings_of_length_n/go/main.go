package main

import (
	"fmt"
)

func getHappyString(n int, k int) string {
	var happyStrings []string
	var dfs func(string)
	dfs = func(currentStr string) {
		if len(currentStr) == n {
			happyStrings = append(happyStrings, currentStr)
			return
		}
		for _, char := range "abc" {
			if len(currentStr) > 0 && currentStr[len(currentStr)-1] == byte(char) {
				continue
			}
			dfs(currentStr + string(char))
		}
	}
	dfs("")
	if len(happyStrings) < k {
		return ""
	}
	return happyStrings[k-1]
}

func main() {
	n := 1
	k := 3
	result := getHappyString(n, k)
	fmt.Printf("Result: %#v\n", result)
	if result == "c" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
