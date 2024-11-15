package main

import (
	"fmt"
)

func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, 100000

	for left < right {
		mid := (left + right) / 2
		count := 0

		for _, quantity := range quantities {
			count += (quantity + mid - 1) / mid
		}

		if count <= n {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	n, quantities := 6, []int{11, 6}
	result := minimizedMaximum(n, quantities)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
