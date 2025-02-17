package main

import (
	"fmt"
)

func numTilePossibilities(tiles string) int {
	count := make([]int, 26)
	for _, tile := range tiles {
		count[tile-'A']++
	}
	var dfs func([]int) int
	dfs = func(count []int) int {
		result := 0
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				result++
				count[i]--
				result += dfs(count)
				count[i]++
			}
		}
		return result
	}
	return dfs(count)
}

func main() {
	tiles := "AAB"
	result := numTilePossibilities(tiles)
	fmt.Printf("Result: %#v\n", result)
	if result == 8 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
