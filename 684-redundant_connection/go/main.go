package main

import (
	"fmt"
)

func findRedundantConnection(edges [][]int) []int {
	var findSetLeader func(int) int
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	findSetLeader = func(vertex int) int {
		if parent[vertex] != vertex {
			parent[vertex] = findSetLeader(parent[vertex])
		}
		return parent[vertex]
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		leaderA := findSetLeader(a)
		leaderB := findSetLeader(b)
		if leaderA == leaderB {
			return []int{a, b}
		}
		parent[leaderA] = leaderB
	}
	return []int{}
}

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	result := findRedundantConnection(edges)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[2 3]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
