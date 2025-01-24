package main

import (
	"fmt"
)

func eventualSafeNodes(graph [][]int) []int {
	totalNodes := len(graph)
	nodeColors := make([]int, totalNodes)
	var dfs func(int) bool
	dfs = func(nodeIndex int) bool {
		if nodeColors[nodeIndex] != 0 {
			return nodeColors[nodeIndex] == 2
		}
		nodeColors[nodeIndex] = 1
		for _, nextNodeIndex := range graph[nodeIndex] {
			if !dfs(nextNodeIndex) {
				return false
			}
		}
		nodeColors[nodeIndex] = 2
		return true
	}
	safeNodes := []int{}
	for nodeIndex := 0; nodeIndex < totalNodes; nodeIndex++ {
		if dfs(nodeIndex) {
			safeNodes = append(safeNodes, nodeIndex)
		}
	}
	return safeNodes
}

func main() {
	graph := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	result := eventualSafeNodes(graph)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[2 4 5 6]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
