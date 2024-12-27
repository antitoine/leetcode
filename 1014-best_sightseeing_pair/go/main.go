package main

import (
	"fmt"
)

func maxScoreSightseeingPair(values []int) int {
	maxScore := 0
	maxValuePlusIndex := values[0]
	for i := 1; i < len(values); i++ {
		maxScore = max(maxScore, values[i]-i+maxValuePlusIndex)
		maxValuePlusIndex = max(maxValuePlusIndex, values[i]+i)
	}
	return maxScore
}

func main() {
	values := []int{8, 1, 5, 2, 6}
	result := maxScoreSightseeingPair(values)
	fmt.Printf("Result: %#v\n", result)
	if result == 11 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
