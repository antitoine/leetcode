package main

import (
	"fmt"
	"slices"
)

func maxCount(banned []int, n int, maxSum int) int {
	slices.Sort(banned)
	var nbInt int
	var sum int
	b := 0
	for i := 1; i <= n; i++ {
		for b < len(banned) && banned[b] < i {
			b++
		}
		if b < len(banned) && banned[b] == i {
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		nbInt++
	}
	return nbInt
}

func main() {
	banned := []int{1, 6, 5}
	n := 5
	maxSum := 6
	result := maxCount(banned, n, maxSum)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
