package main

import (
	"fmt"
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	numVertices := len(edges) + 1
	graph := make([][]int, numVertices)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	timestamps := make([]int, numVertices)
	for i := range timestamps {
		timestamps[i] = numVertices
	}

	var dfsUpdateTimeStamps func(vertex, parent, time int) bool
	dfsUpdateTimeStamps = func(vertex, parent, time int) bool {
		if vertex == 0 {
			timestamps[vertex] = time
			return true
		}
		for _, neighbor := range graph[vertex] {
			if neighbor != parent && dfsUpdateTimeStamps(neighbor, vertex, time+1) {
				timestamps[neighbor] = min(timestamps[neighbor], time+1)
				return true
			}
		}
		return false
	}

	dfsUpdateTimeStamps(bob, -1, 0)
	timestamps[bob] = 0

	maximumProfit := math.MinInt

	var dfsCalculateProfit func(vertex, parent, time, profit int)
	dfsCalculateProfit = func(vertex, parent, time, profit int) {
		if time == timestamps[vertex] {
			profit += amount[vertex] / 2
		} else if time < timestamps[vertex] {
			profit += amount[vertex]
		}

		if len(graph[vertex]) == 1 && graph[vertex][0] == parent {
			maximumProfit = max(maximumProfit, profit)
			return
		}

		for _, neighbor := range graph[vertex] {
			if neighbor != parent {
				dfsCalculateProfit(neighbor, vertex, time+1, profit)
			}
		}
	}
	dfsCalculateProfit(0, -1, 0, 0)

	return maximumProfit
}

func main() {
	edges := [][]int{{0, 1}}
	bob := 1
	amount := []int{-7280, 2350}
	result := mostProfitablePath(edges, bob, amount)
	fmt.Printf("Result: %#v\n", result)
	if result == -7280 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
