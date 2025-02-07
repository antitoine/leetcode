package main

import (
	"fmt"
)

func queryResults(limit int, queries [][]int) []int {
	colors := make(map[int]int)
	balls := make(map[int]int)
	result := make([]int, len(queries))
	for i, query := range queries {
		queryBall := query[0]
		queryColor := query[1]

		oldColor := balls[queryBall]
		if oldColor > 0 {
			colors[oldColor]--
			if colors[oldColor] == 0 {
				delete(colors, oldColor)
			}
		}

		colors[queryColor]++
		balls[queryBall] = queryColor

		result[i] = len(colors)
	}
	return result
}

func main() {
	limit := 4
	queries := [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}
	result := queryResults(limit, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[1 2 2 3]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
