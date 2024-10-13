package main

import (
	"fmt"
	"sort"
)

type Info struct {
	occ int
	pos int
}

func arrayRankTransform(arr []int) []int {
	items := make(map[int]Info)
	sorted := make([]int, 0, len(arr))
	for _, val := range arr {
		if item, found := items[val]; found {
			item.occ++
			items[val] = item
		} else {
			sorted = append(sorted, val)
			items[val] = Info{1, -1}
		}
	}
	sort.Ints(sorted)
	for idx, val := range sorted {
		item := items[val]
		item.pos = idx + 1
		items[val] = item
	}
	result := make([]int, len(arr))
	for idx, val := range arr {
		item := items[val]
		result[idx] = item.pos
	}
	return result
}

func main() {
	input := []int{40, 10, 20, 30}
	result := arrayRankTransform(input)
	fmt.Printf("Result: %#v\n", result)
	if result[0] == 4 && result[1] == 1 && result[2] == 2 && result[3] == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
