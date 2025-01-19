package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	Height int
	X      int
	Y      int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	rows, cols := len(heightMap), len(heightMap[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Add all boundary cells to the heap
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				heap.Push(minHeap, Cell{Height: heightMap[i][j], X: i, Y: j})
				visited[i][j] = true
			}
		}
	}

	trappedWater := 0
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for minHeap.Len() > 0 {
		cell := heap.Pop(minHeap).(Cell)

		for _, dir := range directions {
			nx, ny := cell.X+dir[0], cell.Y+dir[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] {
				trappedWater += max(0, cell.Height-heightMap[nx][ny])
				visited[nx][ny] = true
				heap.Push(minHeap, Cell{Height: max(cell.Height, heightMap[nx][ny]), X: nx, Y: ny})
			}
		}
	}

	return trappedWater
}

func main() {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	result := trapRainWater(heightMap)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
