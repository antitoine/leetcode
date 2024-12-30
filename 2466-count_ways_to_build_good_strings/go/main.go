package main

import "fmt"

const MOD int = 1e9 + 7

func countGoodStrings(low, high, zero, one int) int {
	memo := make(map[int]int)
	var dfs func(current int) int
	dfs = func(current int) int {
		if current > high {
			return 0
		}
		if val, exists := memo[current]; exists {
			return val
		}
		count := 0
		if current >= low && current <= high {
			count = 1
		}
		count = (count + dfs(current+zero) + dfs(current+one)) % MOD
		memo[current] = count
		return count
	}

	return dfs(0)
}

func main() {
	low := 3
	high := 3
	zero := 1
	one := 1
	result := countGoodStrings(low, high, zero, one)
	fmt.Printf("Result: %#v\n", result)
	if result == 8 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
