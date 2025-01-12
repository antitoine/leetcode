package main

import (
	"fmt"
)

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var balance int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || locked[i] == '0' {
			balance++
		} else if balance > 0 {
			balance--
		} else {
			return false
		}
	}
	balance = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			balance++
		} else if balance > 0 {
			balance--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "))()))"
	locked := "010100"
	result := canBeValid(s, locked)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
