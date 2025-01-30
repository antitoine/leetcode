package main

import (
	"container/list"
	"fmt"
)

func magnificentSets(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n+1)

	var componentNodes []int
	var dfs func(int)
	dfs = func(node int) {
		componentNodes = append(componentNodes, node)
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	bfs := func(startNode int) int {
		maximumDistance := 1
		distances := make([]int, n+1)
		for i := range distances {
			distances[i] = 1 << 30 // Simulate infinity
		}
		distances[startNode] = 1

		queue := list.New()
		queue.PushBack(startNode)

		for queue.Len() > 0 {
			currentNode := queue.Remove(queue.Front()).(int)
			for _, neighbor := range graph[currentNode] {
				if distances[neighbor] == 1<<30 { // Unvisited node
					distances[neighbor] = distances[currentNode] + 1
					maximumDistance = distances[neighbor]
					queue.PushBack(neighbor)
				}
			}
		}

		for _, node := range componentNodes {
			if distances[node] == 1<<30 {
				maximumDistance++
				distances[node] = maximumDistance
			}
		}

		for _, node := range componentNodes {
			for _, neighbor := range graph[node] {
				if abs(distances[node]-distances[neighbor]) != 1 {
					return -1
				}
			}
		}
		return maximumDistance
	}

	totalSets := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			componentNodes = []int{}
			dfs(i)

			maxDist := -1
			for _, node := range componentNodes {
				result := bfs(node)
				if result == -1 {
					return -1 // If any component violates conditions, return -1 early
				}
				if result > maxDist {
					maxDist = result
				}
			}

			totalSets += maxDist
		}
	}

	return totalSets
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n := 6
	edges := [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}
	result := magnificentSets(n, edges)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
