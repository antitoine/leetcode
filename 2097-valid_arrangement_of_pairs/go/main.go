package main

import (
	"container/list"
	"fmt"
)

func validArrangement(pairs [][]int) [][]int {
	var ans [][]int
	graph := make(map[int]*list.List)
	outDegree := make(map[int]int)
	inDegree := make(map[int]int)

	// Build the graph and degree counts
	for _, pair := range pairs {
		start, end := pair[0], pair[1]
		if graph[start] == nil {
			graph[start] = list.New()
		}
		graph[start].PushBack(end)
		outDegree[start]++
		inDegree[end]++
	}

	// Find start node
	startNode := getStartNode(graph, outDegree, inDegree, pairs)

	// Perform Eulerian path traversal
	euler(graph, startNode, &ans)

	// Reverse the result to get the path in the correct order
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}

func getStartNode(graph map[int]*list.List, outDegree, inDegree map[int]int, pairs [][]int) int {
	for u := range graph {
		if outDegree[u]-inDegree[u] == 1 {
			return u
		}
	}
	return pairs[0][0] // Arbitrarily choose a node
}

func euler(graph map[int]*list.List, u int, ans *[][]int) {
	stack := graph[u]
	for stack != nil && stack.Len() > 0 {
		e := stack.Front()
		stack.Remove(e)
		v := e.Value.(int)
		euler(graph, v, ans)
		*ans = append(*ans, []int{u, v})
	}
}

func main() {
	pairs := [][]int{
		{5, 1},
		{4, 5},
		{11, 9},
		{9, 4},
	}
	result := validArrangement(pairs)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[[11 9] [9 4] [4 5] [5 1]]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
