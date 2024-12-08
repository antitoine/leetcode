package main

import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	maxValues := make([]int, n)
	maxValues[0] = events[0][2]

	for i := 1; i < n; i++ {
		maxValues[i] = max(maxValues[i-1], events[i][2])
	}

	maxSum := 0
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, events[i][2])

		j := sort.Search(i, func(j int) bool {
			return events[j][1] >= events[i][0]
		}) - 1

		if j >= 0 {
			maxSum = max(maxSum, events[i][2]+maxValues[j])
		}
	}

	return maxSum
}

func main() {
	events := [][]int{
		{1, 3, 2},
		{4, 5, 2},
		{2, 4, 3},
	}
	result := maxTwoEvents(events)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
