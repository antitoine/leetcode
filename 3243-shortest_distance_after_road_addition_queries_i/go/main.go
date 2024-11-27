package main

import (
	"container/list"
	"fmt"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = i
	}

	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		u, v := query[0], query[1]
		graph[u] = append(graph[u], v)
		if dist[u]+1 < dist[v] {
			dist[v] = dist[u] + 1
			bfs(graph, v, dist)
		}
		result[i] = dist[n-1]
	}

	return result
}

func bfs(graph [][]int, start int, dist []int) {
	q := list.New()
	q.PushBack(start)

	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		for _, v := range graph[u] {
			if dist[u]+1 < dist[v] {
				dist[v] = dist[u] + 1
				q.PushBack(v)
			}
		}
	}
}

func main() {
	n, queries := 5, [][]int{{2, 4}, {0, 2}, {0, 4}}
	result := shortestDistanceAfterQueries(n, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[3 2 1]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
