package main

import (
	"fmt"
	"maps"
	"slices"
)

func longestSquareStreak(nums []int) int {
	valuesExplored := make(map[int]bool)
	for _, num := range nums {
		valuesExplored[num] = false
	}
	maxSquareStreakLength := -1
	for _, num := range slices.Sorted(maps.Keys(valuesExplored)) {
		if explored := valuesExplored[num]; explored {
			continue
		}
		currentSquareStreakLength := 1
		nextNum := num * num
		explored, found := valuesExplored[nextNum]
		for found && !explored {
			currentSquareStreakLength++
			nextNum = nextNum * nextNum
			explored, found = valuesExplored[nextNum]
		}
		if currentSquareStreakLength > 1 && currentSquareStreakLength > maxSquareStreakLength {
			maxSquareStreakLength = currentSquareStreakLength
		}
	}
	return maxSquareStreakLength
}

func main() {
	input := []int{4, 3, 6, 16, 8, 2}
	result := longestSquareStreak(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
