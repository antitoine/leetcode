package main

import (
	"fmt"
)

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	ans := 0
	cnt := 0
	for i := 0; i < n*2; i++ {
		if i > 0 && colors[i%n] == colors[(i-1)%n] {
			cnt = 1
		} else {
			cnt++
		}
		if i >= n && cnt >= k {
			ans++
		}
	}
	return ans
}

func main() {
	colors := []int{0, 1, 0, 1, 0}
	k := 3
	result := numberOfAlternatingGroups(colors, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
