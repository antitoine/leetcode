package main

import (
	"fmt"
)

func canChange(start string, target string) bool {
	n := len(start)
	i, j := 0, 0

	for i <= n && j <= n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}
		if i == n || j == n {
			return i == n && j == n
		}
		if start[i] != target[j] {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		i++
		j++
	}

	return true
}

func main() {
	start := "_L__R__R_"
	target := "L______RR"
	result := canChange(start, target)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
