package main

import (
	"fmt"
)

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	diameter1 := getDiameter(edges1)
	diameter2 := getDiameter(edges2)
	combinedDiameter := (diameter1+1)/2 + (diameter2+1)/2 + 1
	return max(diameter1, diameter2, combinedDiameter)
}

func getDiameter(edges [][]int) int {
	n := len(edges) + 1
	graph := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	maxDiameter := 0
	maxDepth(graph, 0, -1, &maxDiameter)
	return maxDiameter
}

func maxDepth(graph [][]int, u, prev int, maxDiameter *int) int {
	maxSubDepth1, maxSubDepth2 := 0, 0

	for _, v := range graph[u] {
		if v == prev {
			continue
		}
		maxSubDepth := maxDepth(graph, v, u, maxDiameter)
		if maxSubDepth > maxSubDepth1 {
			maxSubDepth2 = maxSubDepth1
			maxSubDepth1 = maxSubDepth
		} else if maxSubDepth > maxSubDepth2 {
			maxSubDepth2 = maxSubDepth
		}
	}

	*maxDiameter = max(*maxDiameter, maxSubDepth1+maxSubDepth2)
	return 1 + maxSubDepth1
}

func main() {
	edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	edges2 := [][]int{{0, 1}}
	result := minimumDiameterAfterMerge(edges1, edges2)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
