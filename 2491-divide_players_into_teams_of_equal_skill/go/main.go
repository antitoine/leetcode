package main

import (
	"fmt"
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	startIdx := 0
	endIdx := len(skill) - 1
	sumGroup := skill[startIdx] + skill[endIdx]
	chemistry := int64(skill[startIdx]) * int64(skill[endIdx])
	startIdx++
	endIdx--
	for startIdx < endIdx {
		otherSumGroup := skill[startIdx] + skill[endIdx]
		if otherSumGroup != sumGroup {
			return -1
		}
		chemistry += int64(skill[startIdx]) * int64(skill[endIdx])
		startIdx++
		endIdx--
	}
	return chemistry
}

func main() {
	input := []int{3, 2, 5, 1, 3, 4}
	result := dividePlayers(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 22 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
