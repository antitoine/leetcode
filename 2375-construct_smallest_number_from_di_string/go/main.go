package main

import (
	"fmt"
)

func smallestNumber(pattern string) string {
	visited := make([]bool, 10)
	var result string
	var tempNumber string

	var dfs func(index int)
	dfs = func(index int) {
		if result != "" {
			return
		}

		if index == len(pattern)+1 {
			result = tempNumber
			return
		}

		for i := 1; i < 10; i++ {
			if !visited[i] {
				if index > 0 && pattern[index-1] == 'I' && tempNumber[len(tempNumber)-1]-'0' >= byte(i) {
					continue
				}

				if index > 0 && pattern[index-1] == 'D' && tempNumber[len(tempNumber)-1]-'0' <= byte(i) {
					continue
				}

				visited[i] = true
				tempNumber += fmt.Sprintf("%d", i)
				dfs(index + 1)
				tempNumber = tempNumber[:len(tempNumber)-1]
				visited[i] = false
			}
		}
	}

	dfs(0)
	return result
}

func main() {
	pattern := "IIIDIDDD"
	result := smallestNumber(pattern)
	fmt.Printf("Result: %#v\n", result)
	if result == "123549876" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
