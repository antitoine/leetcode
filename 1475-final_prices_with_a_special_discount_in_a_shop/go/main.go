package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	result := finalPrices(prices)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%#v", result) == "[]int{4, 2, 4, 2, 3}" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
