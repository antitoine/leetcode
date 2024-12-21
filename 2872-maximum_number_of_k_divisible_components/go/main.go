package main

import (
	"fmt"
)

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	ans := 0
	graph := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(u, prev int) int
	dfs = func(u, prev int) int {
		treeSum := values[u]

		for _, v := range graph[u] {
			if v != prev {
				treeSum += dfs(v, u)
			}
		}

		if treeSum%k == 0 {
			ans++
		}
		return treeSum
	}

	dfs(0, -1)
	return ans
}

func main() {
	n, edges, values, k := 5, [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 8, 1, 4, 4}, 6
	result := maxKDivisibleComponents(n, edges, values, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
