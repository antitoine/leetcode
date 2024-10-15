package main

import (
	"fmt"
)

func minimumSteps(s string) int64 {
	var whiteBalls int64
	for _, ball := range s {
		if ball == '0' {
			whiteBalls++
		}
	}
	var score int64
	var stackedBlackBalls int64
	var currentNbWhiteBalls int64
	for _, ball := range s {
		if currentNbWhiteBalls == whiteBalls {
			break
		}
		if ball == '1' {
			stackedBlackBalls++
		} else {
			currentNbWhiteBalls++
			score += stackedBlackBalls
		}
	}
	return score
}

func main() {
	input := "101"
	result := minimumSteps(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
