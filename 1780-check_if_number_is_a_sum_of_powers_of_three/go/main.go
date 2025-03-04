package main

import (
	"fmt"
)

func checkPowersOfThree(n int) bool {
	for n > 0 {
		remainder := n % 3
		if remainder > 1 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	n := 12
	result := checkPowersOfThree(n)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
