package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	x, y, cost int
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minCost(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Point{0, 0, 0})

	visited := make(map[[2]int]bool)

	for pq.Len() > 0 {
		point := heap.Pop(pq).(Point)
		i, j, cost := point.x, point.y, point.cost

		if visited[[2]int{i, j}] {
			continue
		}

		visited[[2]int{i, j}] = true

		if i == rows-1 && j == cols-1 {
			return cost
		}

		for k := 1; k <= 4; k++ {
			x, y := i+directions[k][0], j+directions[k][1]
			if x >= 0 && x < rows && y >= 0 && y < cols {
				if grid[i][j] == k {
					heap.Push(pq, Point{x, y, cost})
				} else {
					heap.Push(pq, Point{x, y, cost + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	grid := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}
	result := minCost(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
