package main

import (
	"fmt"
	"strconv"
)

func canSplitToOriginal(numStr string, startIndex int, targetSum int) bool {
	numLength := len(numStr)
	if startIndex >= numLength {
		return targetSum == 0
	}
	currentNum := 0
	for j := startIndex; j < numLength; j++ {
		currentNum = currentNum*10 + int(numStr[j]-'0')
		if currentNum > targetSum {
			break
		}
		if canSplitToOriginal(numStr, j+1, targetSum-currentNum) {
			return true
		}
	}
	return false
}

func punishmentNumber(n int) int {
	totalSum := 0
	for i := 1; i <= n; i++ {
		squareNum := i * i
		if canSplitToOriginal(strconv.Itoa(squareNum), 0, i) {
			totalSum += squareNum
		}
	}
	return totalSum
}

func main() {
	n := 10
	result := punishmentNumber(n)
	fmt.Printf("Result: %#v\n", result)
	if result == 182 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
