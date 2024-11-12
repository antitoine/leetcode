package main

import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	k := 0
	maxItems := make([][]int, 1, len(items))
	maxItems[k] = items[0]
	for i := 1; i < len(items); i++ {
		if items[i][0] == maxItems[k][0] && maxItems[k][1] < items[i][1] {
			maxItems[k][1] = items[i][1]
		} else if maxItems[k][1] < items[i][1] {
			k++
			maxItems = append(maxItems, items[i])
		}
	}
	for q, price := range queries {
		maxBeauty := 0
		for i := 0; i < len(maxItems) && maxItems[i][0] <= price; i++ {
			if maxBeauty < maxItems[i][1] {
				maxBeauty = maxItems[i][1]
			}
		}
		queries[q] = maxBeauty
	}
	return queries
}

func main() {
	items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	queries := []int{1, 2, 3, 4, 5, 6}
	result := maximumBeauty(items, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[2 4 5 5 6 6]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
