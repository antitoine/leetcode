package main

import (
	"container/list"
	"fmt"
)

func maximumInvitations(favorite []int) int {
	maxCycle := func(favorites []int) int {
		n := len(favorites)
		visited := make([]bool, n)
		maxAns := 0
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			cycle := []int{}
			j := i
			for !visited[j] {
				cycle = append(cycle, j)
				visited[j] = true
				j = favorites[j]
			}
			for k, v := range cycle {
				if v == j {
					maxAns = max(maxAns, len(cycle)-k)
					break
				}
			}
		}
		return maxAns
	}
	topologicalSort := func(favorites []int) int {
		n := len(favorites)
		indegree := make([]int, n)
		longestPath := make([]int, n)
		for i := 0; i < n; i++ {
			indegree[favorites[i]]++
			longestPath[i] = 1
		}
		queue := list.New()
		for i, d := range indegree {
			if d == 0 {
				queue.PushBack(i)
			}
		}
		for queue.Len() > 0 {
			current := queue.Remove(queue.Front()).(int)
			longestPath[favorites[current]] = max(longestPath[favorites[current]], longestPath[current]+1)
			indegree[favorites[current]]--
			if indegree[favorites[current]] == 0 {
				queue.PushBack(favorites[current])
			}
		}
		result := 0
		for i, v := range favorites {
			if i == favorites[v] {
				result += longestPath[i]
			}
		}
		return result
	}
	return max(maxCycle(favorite), topologicalSort(favorite))
}

func main() {
	favorite := []int{2, 2, 1, 2}
	result := maximumInvitations(favorite)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
